package models

import "time"

type Users struct {
	ID        int64
	UserName  string    `binding:"required"`
	Email     string    `binding:"required"`
	Password  string    `binding:"required"`
	Role      string    `binding:"required"`
	CreatedAt time.Time `binding:"required"`
	UpdatedAt time.Time `binding:"required"`
}
