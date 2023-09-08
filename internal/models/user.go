package models

import (
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserInterface interface {
	GetID() int
	GetName() string
	GetApiKey() string
	GetCreatedAt() string
	GetUpdatedAt() string
}

func UserModel(u UserInterface) *User {
	created_at, _ := time.Parse("2006-01-02 15:04:05", u.GetCreatedAt())
	updated_at, _ := time.Parse("2006-01-02 15:04:05", u.GetUpdatedAt())
	return &User{
		Id:        u.GetID(),
		Name:      u.GetName(),
		ApiKey:    u.GetApiKey(),
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}
}
