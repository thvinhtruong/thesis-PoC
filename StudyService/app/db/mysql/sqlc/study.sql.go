// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: study.sql

package db

import (
	"context"
	"database/sql"
)

const createModule = `-- name: CreateModule :execresult
INSERT INTO ` + "`" + `Module` + "`" + ` (` + "`" + `name` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `created_at` + "`" + `, ` + "`" + `updated_at` + "`" + `) VALUES (?, ?, NOW(), NOW())
`

type CreateModuleParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) CreateModule(ctx context.Context, arg CreateModuleParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createModule, arg.Name, arg.Description)
}

const createUserRecord = `-- name: CreateUserRecord :exec
INSERT INTO ` + "`" + `UserRecord` + "`" + ` (` + "`" + `user_id` + "`" + `, ` + "`" + `module_id` + "`" + `, ` + "`" + `weight` + "`" + `, ` + "`" + `score` + "`" + `, ` + "`" + `created_at` + "`" + `, ` + "`" + `updated_at` + "`" + `) VALUES (?, ?, ?, ?, NOW(), NOW())
`

type CreateUserRecordParams struct {
	UserID   int32 `json:"user_id"`
	ModuleID int32 `json:"module_id"`
	Weight   int32 `json:"weight"`
	Score    int32 `json:"score"`
}

func (q *Queries) CreateUserRecord(ctx context.Context, arg CreateUserRecordParams) error {
	_, err := q.db.ExecContext(ctx, createUserRecord,
		arg.UserID,
		arg.ModuleID,
		arg.Weight,
		arg.Score,
	)
	return err
}

const deleteModule = `-- name: DeleteModule :exec
DELETE FROM ` + "`" + `Module` + "`" + ` WHERE ` + "`" + `id` + "`" + ` = ?
`

func (q *Queries) DeleteModule(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteModule, id)
	return err
}

const deleteUserRecord = `-- name: DeleteUserRecord :exec
DELETE FROM ` + "`" + `UserRecord` + "`" + ` WHERE ` + "`" + `id` + "`" + ` = ?
`

func (q *Queries) DeleteUserRecord(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUserRecord, id)
	return err
}

const getModule = `-- name: GetModule :one
SELECT id, name, description, created_at, updated_at FROM ` + "`" + `Module` + "`" + ` WHERE ` + "`" + `id` + "`" + ` = ? FOR UPDATE
`

func (q *Queries) GetModule(ctx context.Context, id int32) (Module, error) {
	row := q.db.QueryRowContext(ctx, getModule, id)
	var i Module
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserRecord = `-- name: GetUserRecord :many
SELECT name, weight, score FROM ` + "`" + `UserRecord` + "`" + ` ur 
    INNER JOIN ` + "`" + `Module` + "`" + ` m
    ON ur.module_id = m.id 
    WHERE ur.user_id = ? 
    ORDER BY ur.id FOR UPDATE
`

type GetUserRecordRow struct {
	Name   string `json:"name"`
	Weight int32  `json:"weight"`
	Score  int32  `json:"score"`
}

func (q *Queries) GetUserRecord(ctx context.Context, userID int32) ([]GetUserRecordRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserRecord, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserRecordRow
	for rows.Next() {
		var i GetUserRecordRow
		if err := rows.Scan(&i.Name, &i.Weight, &i.Score); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserRecordById = `-- name: GetUserRecordById :one
SELECT id, user_id, module_id, weight, score, created_at, updated_at FROM ` + "`" + `UserRecord` + "`" + ` WHERE ` + "`" + `id` + "`" + ` = ? FOR UPDATE
`

func (q *Queries) GetUserRecordById(ctx context.Context, id int32) (Userrecord, error) {
	row := q.db.QueryRowContext(ctx, getUserRecordById, id)
	var i Userrecord
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ModuleID,
		&i.Weight,
		&i.Score,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserRecordByUserIdAndModuleId = `-- name: GetUserRecordByUserIdAndModuleId :one
SELECT id, user_id, module_id, weight, score, created_at, updated_at FROM ` + "`" + `UserRecord` + "`" + ` WHERE ` + "`" + `user_id` + "`" + ` = ? AND ` + "`" + `module_id` + "`" + ` = ? FOR UPDATE
`

type GetUserRecordByUserIdAndModuleIdParams struct {
	UserID   int32 `json:"user_id"`
	ModuleID int32 `json:"module_id"`
}

func (q *Queries) GetUserRecordByUserIdAndModuleId(ctx context.Context, arg GetUserRecordByUserIdAndModuleIdParams) (Userrecord, error) {
	row := q.db.QueryRowContext(ctx, getUserRecordByUserIdAndModuleId, arg.UserID, arg.ModuleID)
	var i Userrecord
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ModuleID,
		&i.Weight,
		&i.Score,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listModules = `-- name: ListModules :many
SELECT id, name, description, created_at, updated_at FROM ` + "`" + `Module` + "`" + ` WHERE ` + "`" + `id` + "`" + ` > ? ORDER BY ` + "`" + `id` + "`" + ` LIMIT ?
`

type ListModulesParams struct {
	ID    int32 `json:"id"`
	Limit int32 `json:"limit"`
}

func (q *Queries) ListModules(ctx context.Context, arg ListModulesParams) ([]Module, error) {
	rows, err := q.db.QueryContext(ctx, listModules, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Module
	for rows.Next() {
		var i Module
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserRecords = `-- name: ListUserRecords :many
SELECT id, user_id, module_id, weight, score, created_at, updated_at FROM ` + "`" + `UserRecord` + "`" + ` WHERE ` + "`" + `id` + "`" + ` > ? ORDER BY ` + "`" + `id` + "`" + ` LIMIT ?
`

type ListUserRecordsParams struct {
	ID    int32 `json:"id"`
	Limit int32 `json:"limit"`
}

func (q *Queries) ListUserRecords(ctx context.Context, arg ListUserRecordsParams) ([]Userrecord, error) {
	rows, err := q.db.QueryContext(ctx, listUserRecords, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Userrecord
	for rows.Next() {
		var i Userrecord
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ModuleID,
			&i.Weight,
			&i.Score,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateModule = `-- name: UpdateModule :exec
UPDATE ` + "`" + `Module` + "`" + ` SET ` + "`" + `name` + "`" + ` = ?, ` + "`" + `description` + "`" + ` = ?, ` + "`" + `updated_at` + "`" + ` = NOW() WHERE ` + "`" + `id` + "`" + ` = ?
`

type UpdateModuleParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	ID          int32          `json:"id"`
}

func (q *Queries) UpdateModule(ctx context.Context, arg UpdateModuleParams) error {
	_, err := q.db.ExecContext(ctx, updateModule, arg.Name, arg.Description, arg.ID)
	return err
}

const updateUserRecord = `-- name: UpdateUserRecord :exec
UPDATE ` + "`" + `UserRecord` + "`" + ` SET ` + "`" + `user_id` + "`" + ` = ?, ` + "`" + `module_id` + "`" + ` = ?, ` + "`" + `weight` + "`" + ` = ?, ` + "`" + `score` + "`" + ` = ?, ` + "`" + `updated_at` + "`" + ` = NOW() WHERE ` + "`" + `id` + "`" + ` = ?
`

type UpdateUserRecordParams struct {
	UserID   int32 `json:"user_id"`
	ModuleID int32 `json:"module_id"`
	Weight   int32 `json:"weight"`
	Score    int32 `json:"score"`
	ID       int32 `json:"id"`
}

func (q *Queries) UpdateUserRecord(ctx context.Context, arg UpdateUserRecordParams) error {
	_, err := q.db.ExecContext(ctx, updateUserRecord,
		arg.UserID,
		arg.ModuleID,
		arg.Weight,
		arg.Score,
		arg.ID,
	)
	return err
}

const updateUserRecordScore = `-- name: UpdateUserRecordScore :exec
UPDATE ` + "`" + `UserRecord` + "`" + ` SET ` + "`" + `score` + "`" + ` = ?, ` + "`" + `updated_at` + "`" + ` = NOW() WHERE ` + "`" + `user_id` + "`" + ` = ? AND ` + "`" + `module_id` + "`" + ` = ?
`

type UpdateUserRecordScoreParams struct {
	Score    int32 `json:"score"`
	UserID   int32 `json:"user_id"`
	ModuleID int32 `json:"module_id"`
}

func (q *Queries) UpdateUserRecordScore(ctx context.Context, arg UpdateUserRecordScoreParams) error {
	_, err := q.db.ExecContext(ctx, updateUserRecordScore, arg.Score, arg.UserID, arg.ModuleID)
	return err
}
