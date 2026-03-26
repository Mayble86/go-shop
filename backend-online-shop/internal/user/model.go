package user

import (
	"time"
)

type Role struct {
	ID   int
	Name string
}

type User struct {
	ID        int
	Email     string
	Password  string
	RoleID    int
	CreatedAt time.Time
}
