package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"olaf/internal/model"
	"strings"
)

type userRepo struct {
	pool *pgxpool.Pool
}

func NewUsersPool(pool *pgxpool.Pool) Users {
	return &userRepo{pool: pool}
}

func (u userRepo) AddUser(ctx context.Context, user model.User) (int, error) {
	var id = 0
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return id, fmt.Errorf("AddUser: %w", err)
	}
	email := strings.ToLower(user.Email)
	dbReq := "INSERT INTO users (name, password, email, isActive, group_id) " +
		"VALUES ($1, $2, $3, $4, $5) " +
		"RETURNING id"
	err = u.pool.QueryRow(ctx, dbReq, user.Name, hash, email, user.IsActive, user.Group).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("AddUser: %w", err)
	}
	return id, nil
}

func (u userRepo) EditUser(ctx context.Context, user model.User) error {
	dbReq := "UPDATE USERS SET "
	if user.Name != "" {
		dbReq = fmt.Sprintf("%s name='%s', ", dbReq, user.Name)
	}
	if user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			return fmt.Errorf("EditUser: %w", err)
		}
		dbReq = fmt.Sprintf("%s password='%s', ", dbReq, hash)
	}
	if user.Email != "" {
		email := strings.ToLower(user.Email)
		dbReq = fmt.Sprintf("%s email='%s', ", dbReq, email)
	}
	if user.Group != 0 {
		dbReq = fmt.Sprintf("%s groip_id=%d, ", dbReq, user.Group)
	}
	dbReq = fmt.Sprintf("%s updated=CURRENT_TIMESTAMP WHERE id=%d AND isactive=true", dbReq, user.ID)
	_, err := u.pool.Query(ctx, dbReq)
	if err != nil {
		return fmt.Errorf("EditUser: %w", err)
	}
	return nil
}

func (u userRepo) ListAllUsers(ctx context.Context) ([]model.User, error) {
	users := make([]model.User, 0)
	dbReq := "SELECT id, name, email, isactive, group_id FROM users WHERE isactive=true"
	rows, err := u.pool.Query(ctx, dbReq)
	defer rows.Close()
	if err != nil {
		return users, fmt.Errorf("ListAllUsers: %w", err)
	}
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.IsActive, &user.Group)
		if err != nil {
			return users, fmt.Errorf("ListAllUsers: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (u userRepo) DeleteUserbyID(ctx context.Context, id int) error {
	dbReq := fmt.Sprintf("UPDATE users SET isactive=false, updated=CURRENT_TIMESTAMP WHERE id=%d", id)
	_, err := u.pool.Query(ctx, dbReq)
	if (err != nil) && (err != pgx.ErrNoRows) {
		return fmt.Errorf("DeleteUser: %w", err)
	}
	return nil
}

func (u userRepo) GetUserIDbyEmail(ctx context.Context, email string) (int, error) {
	dbReq := fmt.Sprintf("SELECT id FROM users WHERE email='%s' AND isactive=true", strings.ToLower(email))
	var id int
	err := u.pool.QueryRow(ctx, dbReq).Scan(&id)
	if (err != nil) && (err != pgx.ErrNoRows) {
		return 0, fmt.Errorf("GetUserIDbyEmail: %w", err)
	}
	return id, nil
}
