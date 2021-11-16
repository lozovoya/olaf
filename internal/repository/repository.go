package repository

import (
	"context"
	"olaf/internal/model"
)

type Users interface {
	AddUser(context.Context, model.User) (int, error)
	EditUser(context.Context, model.User) error
	ListAllUsers(context.Context) ([]model.User, error)
	DeleteUserbyID(context.Context, int) error
	GetUserIDbyEmail(ctx context.Context, email string) (int, error)
}

type Groups interface {
	AddGroup(context.Context, model.Group) (int, error)
	EditGroup(context.Context, model.Group) error
	ListAllGroups(context.Context) ([]model.Group, error)
	DeleteGroupbyID(context.Context, int) error
	GetGroupIDbyName(context.Context, string) (int, error)
	GetGroupMembersbyID(context.Context, int) ([]model.User, error)
}
