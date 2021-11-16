package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/suite"
	"olaf/internal/model"
	"testing"
)

const testDSN = "postgres://app:pass@localhost:5432/olaftestdb"

type GroupsTestSuite struct {
	suite.Suite
	testRepo groupsRepo
}

func (suite *GroupsTestSuite) SetupTest() {
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
		"VALUES ('user1', 'password', 'user1@user.com', true, 1); "
	_, err = suite.testRepo.pool.Query(context.Background(), addUserReq)
	if err != nil {
		suite.Error(err)
		suite.Fail("setup failed")
		return
	}
}

func (suite *GroupsTestSuite) TearDownTest() {
	fmt.Println("cleaning up")
	var err error
	_, err = suite.testRepo.pool.Query(context.Background(), "DROP TABLE users, groups CASCADE;")
	if err != nil {
		suite.Error(err)
	}
}

func Test_GroupsSuite(t *testing.T) {
	suite.Run(t, new(GroupsTestSuite))
}

func (suite *GroupsTestSuite) Test_groupsRepo_AddGroup() {

	type args struct {
		ctx   context.Context
		group model.Group
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "add new group",
			args: args{
				ctx: context.Background(),
				group: model.Group{
					Name:     "group2",
					IsActive: true,
				},
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "add wrong group",
			args: args{
				ctx: context.Background(),
				group: model.Group{
					Name:     "group2",
					IsActive: true,
				},
			},
			//want:    2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			got, err := suite.testRepo.AddGroup(tt.args.ctx, tt.args.group)
			if (err != nil) != tt.wantErr {
				fmt.Printf("AddGroup() error = %v, wantErr %v", err, tt.wantErr)
				suite.Fail("test failed")
				return
			}
			if !suite.Equal(tt.want, got) {
				fmt.Printf("AddGroup() error = %v, wantErr %v", err, tt.want)
				suite.Fail("test failed")
			}
		})
	}
}

func (suite *GroupsTestSuite) Test_groupsRepo_DeleteGroupbyID() {
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
			name: "delete existing group",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantErr: false,
		},
		{
			name: "delete non-existing group",
			args: args{
				ctx: context.Background(),
				id:  2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			err := suite.testRepo.DeleteGroupbyID(tt.args.ctx, tt.args.id)
			if (err != nil) && (tt.wantErr == false) {
				fmt.Printf("DeleteGroupbyID() error = %v, wantErr %v", err, tt.wantErr)
				suite.Fail("test failed")
				return
			}
		})
	}
}

func (suite *GroupsTestSuite) Test_groupsRepo_EditGroup() {

	type args struct {
		ctx   context.Context
		group model.Group
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "edit existing group",
			args: args{
				ctx: context.Background(),
				group: model.Group{
					ID:   1,
					Name: "new name for group",
				},
			},
			wantErr: false,
		},
		{
			name: "edit non-existing group",
			args: args{
				ctx: context.Background(),
				group: model.Group{
					ID:   2,
					Name: "new name for group",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			err := suite.testRepo.EditGroup(tt.args.ctx, tt.args.group)
			if (err != nil) && (tt.wantErr == false) {
				fmt.Printf("EditGroup() error = %v, wantErr %v", err, tt.wantErr)
				suite.Fail("test failed")
				return
			}
		})
	}
}

func (suite *GroupsTestSuite) Test_groupsRepo_GetGroupIDbyName() {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "get existing group",
			args: args{
				ctx:  context.Background(),
				name: "group1",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "get non-existing group",
			args: args{
				ctx:  context.Background(),
				name: "group2",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			got, err := suite.testRepo.GetGroupIDbyName(tt.args.ctx, tt.args.name)
			if (err != nil) && (tt.wantErr == false) {
				fmt.Printf("GetGroupIDbyName() error = %v, wantErr %v", err, tt.wantErr)
				suite.Fail("test failed")
				return
			}
			if !suite.Equal(tt.want, got) {
				fmt.Printf("GetGroupIDbyName() got = %v, want %v", got, tt.want)
				suite.Fail("test failed")
			}
		})
	}
}

func (suite *GroupsTestSuite) Test_groupsRepo_GetGroupMembersbyID() {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    []model.User
		wantErr bool
	}{
		{
			name: "get existing group",
			args: args{
				ctx: context.Background(),
				id:  1,
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
			},
			wantErr: false,
		},
		{
			name: "get non-existing group",
			args: args{
				ctx: context.Background(),
				id:  2,
			},
			want:    []model.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			got, err := suite.testRepo.GetGroupMembersbyID(tt.args.ctx, tt.args.id)
			if (err != nil) && (tt.wantErr == false) {
				fmt.Printf("GetGroupMembersbyID() error = %v, wantErr %v", err, tt.wantErr)
				suite.Fail("test failed")
				return
			}
			if !suite.Equal(tt.want, got) {
				fmt.Printf("GetGroupMembersbyID() got = %v, want %v", got, tt.want)
				suite.Fail("test failed")
			}
		})
	}
}

func (suite *GroupsTestSuite) Test_groupsRepo_ListAllGroups() {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Group
		wantErr bool
	}{
		{
			name: "list groups",
			args: args{
				ctx: context.Background(),
			},
			want: []model.Group{
				{
					ID:       1,
					Name:     "group1",
					IsActive: true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			got, err := suite.testRepo.ListAllGroups(tt.args.ctx)
			if (err != nil) && (tt.wantErr == false) {
				fmt.Printf("ListAllGroups() error = %v, wantErr %v", err, tt.wantErr)
				suite.Fail("test failed")
				return
			}
			if !suite.Equal(tt.want, got) {
				fmt.Printf("ListAllGroups() got = %v, want %v", got, tt.want)
				suite.Fail("test failed")
			}
		})
	}
}
