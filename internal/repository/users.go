package repository

import (
	"context"
	"fmt"
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
	dbReq = fmt.Sprintf("%s isactive=%t, updated=CURRENT_TIMESTAMP WHERE id=%d", dbReq, user.IsActive, user.ID)
	_, err := u.pool.Query(ctx, dbReq)
	if err != nil {
		return fmt.Errorf("EditUser: %w", err)
	}
	return nil
}

func (u userRepo) ListAllUsers(ctx context.Context) ([]model.User, error) {
	users := make([]model.User, 0)
	dbReq := "SELECT id, name, email, isactive, group FROM users"
	rows, err := u.pool.Query(ctx, dbReq)
	if err != nil {
		return users, fmt.Errorf("ListAllUsers: %w", err)
	}
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.IsActive, &user.Group)
		users = append(users, user)
	}
	return users, nil
}

func (u userRepo) DeleteUser(ctx context.Context, user model.User) error {
	panic("implement me")
}

func (u userRepo) GetUser(ctx context.Context, email string) (model.User, error) {
	panic("implement me")
}
