-- name: CreateUser :execresult
INSERT INTO User (fullname, email, gender, datecreated)
VALUES(?, ? , ?, NOW());

-- name: CreateUserPassword :exec
INSERT INTO User_Password (user_id, password) VALUES(?, ?);

-- name: ListUsers :many
SELECT * FROM User ORDER BY id FOR UPDATE;

-- name: GetUserByEmail :one
SELECT * FROM User WHERE email = ? LIMIT 1 FOR UPDATE;

-- name: GetUserPassword :one
SELECT * FROM User_Password WHERE user_id = ? LIMIT 1;

-- name: UpdateUserInfo :exec
UPDATE User SET fullname = ?, email = ?, gender = ?, dateupdated = NOW()
WHERE id = ?;

-- name: UpdateUserPassword :exec
UPDATE User_Password SET password = ? WHERE user_id = ?;

-- name: DeleteUser :exec
DELETE FROM User WHERE id = ?;

-- name: GetUser :one
SELECT * FROM User WHERE id = ? LIMIT 1 FOR UPDATE;

-- name: DeleteUserPassword :exec
DELETE FROM User_Password WHERE user_id = ?;
