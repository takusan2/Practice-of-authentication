package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID           int64     `bun:"id,pk,autoincrement"`
	Name         string    `bun:"name,notnull"`
	Email        string    `bun:"email,notnull,unique"`
	PasswordHash string    `bun:"password_hash,notnull"`
	CreatedAt    time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt    time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
