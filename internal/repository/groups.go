package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
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
		return id, fmt.Errorf("AddGroup: %w", err)
	}
	return id, nil
}

func (g groupsRepo) EditGroup(ctx context.Context, group model.Group) error {
	dbReq := "UPDATE groups SET "
	if group.Name != "" {
		dbReq = fmt.Sprintf("%s name='%s', updated=CURRENT_TIMESTAMP WHERE id=%d AND isactive=true",
			dbReq,
			group.Name,
			group.ID)
	}
	_, err := g.pool.Query(ctx, dbReq)
	if err != nil {
		return fmt.Errorf("EditGroup: %w", err)
	}
	return nil
}

func (g groupsRepo) ListAllGroups(ctx context.Context) ([]model.Group, error) {
	var groups = make([]model.Group, 0)
	dbReq := "SELECT id, name, isactive FROM groups WHERE isactive=true"
	rows, err := g.pool.Query(ctx, dbReq)
	defer rows.Close()
	if (err != nil) && (err != pgx.ErrNoRows) {
		return groups, fmt.Errorf("ListAllGroups: %w", err)
	}
	for rows.Next() {
		var group model.Group
		err = rows.Scan(&group.ID, &group.Name, &group.IsActive)
		if err != nil {
			return groups, fmt.Errorf("ListAllGroups: %w", err)
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (g groupsRepo) DeleteGroupbyID(ctx context.Context, id int) error {
	dbReq := fmt.Sprintf("UPDATE groups SET isactive=false, updated=CURRENT_TIMESTAMP WHERE id=%d", id)
	_, err := g.pool.Query(ctx, dbReq)
	if (err != nil) && (err != pgx.ErrNoRows) {
		return fmt.Errorf("DeleteGroupbyID: %w", err)
	}
	return nil
}

func (g groupsRepo) GetGroupIDbyName(ctx context.Context, name string) (int, error) {
	dbReq := fmt.Sprintf("SELECT id FROM groups WHERE name='%s' AND isactive=true", name)
	var id int = 0
	err := g.pool.QueryRow(ctx, dbReq).Scan(&id)
	if (err != nil) && (err != pgx.ErrNoRows) {
		return id, fmt.Errorf("GetGroupIDbyName: %w", err)
	}
	return id, nil
}

func (g groupsRepo) GetGroupMembersbyID(ctx context.Context, id int) ([]model.User, error) {
	dbReq := "SELECT DISTINCT users.id, users.name, users.email, users.isactive, users.group_id " +
		"FROM users, groups " +
		"WHERE (users.group_id=$1 AND users.isactive=true) " +
		"AND (groups.id=$1 AND groups.isactive=true)"
	var users = make([]model.User, 0)
	rows, err := g.pool.Query(ctx, dbReq, id)
	defer rows.Close()
	if (err != nil) && (err != pgx.ErrNoRows) {
		return users, fmt.Errorf("GetGroupMembersbyID: %w", err)
	}
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.IsActive, &user.Group)
		if err != nil {
			return users, fmt.Errorf("GetGroupMembersbyID: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}
