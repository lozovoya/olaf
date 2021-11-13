package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"olaf/internal/model"
)

type groupsRepo struct {
	pool *pgxpool.Pool
}

func NewGroupsRepo(pool *pgxpool.Pool) Groups {
	return &groupsRepo{pool: pool}
}

func (g groupsRepo) AddGroup(ctx context.Context, group model.Group) (int, error) {
	var id = 0
	dbReq := "INSERT INTO groups (name, isactive) " +
		"VALUES ($1, $2) " +
		"RETURNING id"
	err := g.pool.QueryRow(ctx, dbReq, group.Name, group.IsActive).Scan(&id)
	if err != nil {
		log.Println(err)
		return id, fmt.Errorf("AddGroup: %w", err)
	}
	return id, nil
}

func (g groupsRepo) EditGroup(ctx context.Context, group model.Group) (model.Group, error) {
	panic("implement me")
}

func (g groupsRepo) ListAllGroups(ctx context.Context, group model.Group) ([]model.Group, error) {
	panic("implement me")
}

func (g groupsRepo) DeleteGroup(ctx context.Context, group model.Group) error {
	panic("implement me")
}

func (g groupsRepo) GetGroup(ctx context.Context, name string) (model.Group, error) {
	panic("implement me")
}

func (g groupsRepo) GetGroupMembers(ctx context.Context, name string) ([]model.User, error) {
	panic("implement me")
}
