package repository

import (
	"context"
	"olaf/internal/model"
)

type Users interface {
	AddUser (context.Context, ) (model.User, error)
	EditUser (context.Context, model.User)(model.User, error)
	ListAllUsers (context.Context)([]model.User, error)
	DeleteUser (context.Context, model.User) error
	GetUser (ctx context.Context, email string) (model.User, error)
}

type Groups interface {
	AddGroup (context.Context, model.Group) (model.Group, error)
	EditGroup (context.Context, model.Group) (model.Group, error)
	ListAllGroups (context.Context, model.Group) ([]model.Group, error)
	DeleteGroup (context.Context, model.Group) error
	GetGroup (ctx context.Context, name string) (model.Group, error)
	GetGroupMembers (ctx context.Context, name string) ([]model.User, error)
}