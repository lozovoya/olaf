package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/suite"
	"olaf/internal/model"
	"testing"
)

type UsersTestSuite struct {
	suite.Suite
	testRepo userRepo
}

func (suite *UsersTestSuite) SetupTest() {
	fmt.Println("start setup")
	var err error
	suite.testRepo.pool, err = pgxpool.Connect(context.Background(), testDSN)
	if err != nil {
		suite.Error(err)
		suite.Fail("setup failed")
		return
	}

	createTableGroupsReq := "CREATE TABLE groups ( " +
		"id BIGSERIAL PRIMARY KEY," +
		"name VARCHAR(50) NOT NULL UNIQUE, " +
		"isActive BOOLEAN DEFAULT false NOT NULL, " +
		"created  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, " +
		"updated  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP);"
	_, err = suite.testRepo.pool.Query(context.Background(), createTableGroupsReq)
	if err != nil {
		suite.Error(err)
		suite.Fail("setup failed")
		return
	}
	createTableUsersReq := "CREATE TABLE users ( " +
		"id BIGSERIAL PRIMARY KEY, " +
		"name VARCHAR(200) NOT NULL, " +
		"password VARCHAR(200) NOT NULL, " +
		"email VARCHAR (100) NOT NULL UNIQUE, " +
		"isActive BOOLEAN DEFAULT false NOT NULL, " +
		"group_id BIGINT NOT NULL REFERENCES groups, " +
		"created  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, " +
		"updated  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP);"
	_, err = suite.testRepo.pool.Query(context.Background(), createTableUsersReq)
	if err != nil {
		suite.Error(err)
		suite.Fail("setup failed")
		return
	}
	addGroupReq := "INSERT INTO groups (name, isactive) " +
		"VALUES ('group1', true); "
	_, err = suite.testRepo.pool.Query(context.Background(), addGroupReq)
	if err != nil {
		suite.Error(err)
		suite.Fail("setup failed")
		return
	}
	addUserReq := "INSERT INTO users (name, password, email, isActive, group_id) " +
		"VALUES ('user1', 'password', 'user1@user.com', true, 1), " +
		"('user2', 'password', 'user2@user.com', true, 1); "
	_, err = suite.testRepo.pool.Query(context.Background(), addUserReq)
	if err != nil {
		suite.Error(err)
		suite.Fail("setup failed")
		return
	}
}

func (suite *UsersTestSuite) TearDownTest() {
	fmt.Println("cleaning up")
	var err error
	_, err = suite.testRepo.pool.Query(context.Background(), "DROP TABLE users, groups CASCADE;")
	if err != nil {
		suite.Error(err)
	}
}

func Test_UsersSuite(t *testing.T) {
	suite.Run(t, new(UsersTestSuite))
}

func (suite *UsersTestSuite) Test_userRepo_AddUser() {

	type args struct {
		ctx  context.Context
		user model.User
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "add new user",
			args: args{
				ctx: context.Background(),
				user: model.User{
					Name:     "user3",
					Password: "password",
					Email:    "user3@user.com",
					IsActive: true,
					Group:    1,
				},
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "add existing user",
			args: args{
				ctx: context.Background(),
				user: model.User{
					Name:     "user3",
					Password: "password",
					Email:    "user3@user.com",
					IsActive: true,
					Group:    1,
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "add new user to unexisting group",
			args: args{
				ctx: context.Background(),
				user: model.User{
					Name:     "user4",
					Password: "password",
					Email:    "user4@user.com",
					IsActive: true,
					Group:    4,
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			got, err := suite.testRepo.AddUser(tt.args.ctx, tt.args.user)
			if (err != nil) && (tt.wantErr == false) {
				fmt.Printf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
				suite.Fail("test failed")
				return
			}
			if got != tt.want {
				fmt.Printf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
				suite.Fail("test failed")
			}
		})
	}
}

func (suite *UsersTestSuite) Test_userRepo_DeleteUserbyID() {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete existing user",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantErr: false,
		},
		{
			name: "delete non-existing user",
			args: args{
				ctx: context.Background(),
				id:  3,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			err := suite.testRepo.DeleteUserbyID(tt.args.ctx, tt.args.id)
			if (err != nil) && (tt.wantErr == false) {
				fmt.Printf("DeleteUserbyID() error = %v, wantErr %v", err, tt.wantErr)
				suite.Fail("test failed")
			}
		})
	}
}

func (suite *UsersTestSuite) Test_userRepo_EditUser() {
	type args struct {
		ctx  context.Context
		user model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "edit existing user",
			args: args{
				ctx: context.Background(),
				user: model.User{
					ID:       1,
					Name:     "new name for user",
					Password: "new password",
					Email:    "",
					IsActive: true,
					Group:    1,
				},
			},
			wantErr: false,
		},
		{
			name: "edit non-existing user",
			args: args{
				ctx: context.Background(),
				user: model.User{
					ID:       3,
					Name:     "new name for user",
					Password: "new password",
					Email:    "",
					IsActive: true,
					Group:    1,
				},
			},
			wantErr: true,
		},
		{
			name: "duplicate email",
			args: args{
				ctx: context.Background(),
				user: model.User{
					ID:       1,
					Name:     "new name for user",
					Password: "new password",
					Email:    "user2@user.com",
					IsActive: true,
					Group:    1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			err := suite.testRepo.EditUser(tt.args.ctx, tt.args.user)
			if (err != nil) && (tt.wantErr == false) {
				fmt.Printf("EditUser() error = %v, wantErr %v", err, tt.wantErr)
				suite.Fail("test failed")
			}
		})
	}
}

func (suite *UsersTestSuite) Test_userRepo_GetUserIDbyEmail() {
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "get existing user",
			args: args{
				ctx:   context.Background(),
				email: "user1@user.com",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "get non-existing user",
			args: args{
				ctx:   context.Background(),
				email: "user3@user.com",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			got, err := suite.testRepo.GetUserIDbyEmail(tt.args.ctx, tt.args.email)
			if (err != nil) && (tt.wantErr == false) {
				fmt.Printf("GetUserIDbyEmail() error = %v, wantErr %v", err, tt.wantErr)
				suite.Fail("test failed")
				return
			}
			if got != tt.want {
				fmt.Printf("GetUserIDbyEmail() got = %v, want %v", got, tt.want)
				suite.Fail("test failed")
			}
		})
	}
}

func (suite *UsersTestSuite) Test_userRepo_ListAllUsers() {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []model.User
		wantErr bool
	}{
		{
			name: "list all users",
			args: args{
				ctx: context.Background(),
			},
			want: []model.User{
				{
					ID:       1,
					Name:     "user1",
					Password: "",
					Email:    "user1@user.com",
					IsActive: true,
					Group:    1,
				},
				{
					ID:       2,
					Name:     "user2",
					Password: "",
					Email:    "user2@user.com",
					IsActive: true,
					Group:    1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			got, err := suite.testRepo.ListAllUsers(tt.args.ctx)
			if (err != nil) && (tt.wantErr == false) {
				fmt.Printf("ListAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				suite.Fail("test failed")
				return
			}
			if !suite.Equal(tt.want, got) {
				fmt.Printf("ListAllUsers() got = %v, want %v", got, tt.want)
				suite.Fail("test failed")
			}
		})
	}
}
