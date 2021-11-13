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
		"VALUES $1, $2, $3, $4, $5 " +
		"RETURNING id"
	err = u.pool.QueryRow(ctx, dbReq, user.Name, hash, email, user.IsActive, user.Group).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("AddUser: %w", err)
	}
	return id, nil
}

func (u userRepo) EditUser(ctx context.Context, user model.User) (model.User, error) {
	panic("implement me")
}

func (u userRepo) ListAllUsers(ctx context.Context) ([]model.User, error) {
	panic("implement me")
}

func (u userRepo) DeleteUser(ctx context.Context, user model.User) error {
	panic("implement me")
}

func (u userRepo) GetUser(ctx context.Context, email string) (model.User, error) {
	panic("implement me")
}
