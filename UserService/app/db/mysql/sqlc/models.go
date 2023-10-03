// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"
)

type User struct {
	ID          int32        `json:"id"`
	Fullname    string       `json:"fullname"`
	Gender      string       `json:"gender"`
	Email       string       `json:"email"`
	Datecreated time.Time    `json:"datecreated"`
	Dateupdated sql.NullTime `json:"dateupdated"`
}

type UserPassword struct {
	ID       int32  `json:"id"`
	UserID   int32  `json:"user_id"`
	Password string `json:"password"`
}
