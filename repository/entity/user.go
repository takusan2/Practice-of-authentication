package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID           int64     `bun:"id,pk,autoincrement"`
	Name         string    `bun:"name,notnull"`
	Email        string    `bun:"email,notnull"`
	PasswordHash string    `bun:"password_hash,notnull"`
	Salt         string    `bun:"salt,notnull"`
	CreatedAt    time.Time `bun:"created_at,notnull"`
}
