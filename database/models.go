// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	AuthID     string
	Username   string
	Email      string
	PictureUrl sql.NullString
}
