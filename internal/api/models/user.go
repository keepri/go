package models

import (
	"time"

	"github.com/keepri/go/internal/database"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UserModel(u *database.User) *User {
	created_at, _ := time.Parse("2006-01-02 15:04:05", u.CreatedAt)
	updated_at, _ := time.Parse("2006-01-02 15:04:05", u.UpdatedAt)

	return &User{
		Id:        int(u.ID),
		Name:      u.Name,
		ApiKey:    u.ApiKey,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}
}
