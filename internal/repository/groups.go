package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"olaf/internal/model"
)

type groupsRepo struct {
	pool *pgxpool.Pool
}

func NewGroupsRepo(pool *pgxpool.Pool) Groups {
	return &groupsRepo{pool: pool}
}

func (g groupsRepo) AddGroup(ctx context.Context, group model.Group) (model.Group, error) {
	panic("implement me")
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
