-- name: CreateModule :execresult
INSERT INTO `Module` (`name`, `description`, `created_at`, `updated_at`) VALUES (?, ?, NOW(), NOW());

-- name: GetModule :one
SELECT * FROM `Module` WHERE `id` = ? FOR UPDATE;

-- name: UpdateModule :exec
UPDATE `Module` SET `name` = ?, `description` = ?, `updated_at` = NOW() WHERE `id` = ?;

-- name: DeleteModule :exec
DELETE FROM `Module` WHERE `id` = ?;

-- name: ListModules :many
SELECT * FROM `Module` WHERE `id` > ? ORDER BY `id` LIMIT ?;

-- name: CreateUserRecord :exec
INSERT INTO `UserRecord` (`user_id`, `module_id`, `weight`, `score`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, NOW(), NOW());

-- name: GetUserRecordById :one
SELECT * FROM `UserRecord` WHERE `id` = ? FOR UPDATE;

-- name: GetUserRecordByUserIdAndModuleId :one
SELECT * FROM `UserRecord` WHERE `user_id` = ? AND `module_id` = ? FOR UPDATE;

-- name: UpdateUserRecord :exec
UPDATE `UserRecord` SET `user_id` = ?, `module_id` = ?, `weight` = ?, `score` = ?, `updated_at` = NOW() WHERE `id` = ?;

-- name: UpdateUserRecordScore :exec
UPDATE `UserRecord` SET `score` = ?, `updated_at` = NOW() WHERE `user_id` = ? AND `module_id` = ?;

-- name: DeleteUserRecord :exec
DELETE FROM `UserRecord` WHERE `id` = ?;

-- name: ListUserRecords :many
SELECT * FROM `UserRecord` WHERE `id` > ? ORDER BY `id` LIMIT ?;

-- name: GetUserRecord :many
SELECT name, weight, score FROM `UserRecord` ur 
    INNER JOIN `Module` m
    ON ur.module_id = m.id 
    WHERE ur.user_id = ? 
    ORDER BY ur.id FOR UPDATE;