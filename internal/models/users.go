package models

import (
	"blog/root/internal/database"
	"errors"
	"fmt"
	"time"
)

type Credentials struct {
	ID       int64
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

type Users struct {
	ID        int64
	UserName  string    `json:"user_name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Role      string    `json:"role" binding:"required,oneof=author admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *Users) Create() error {
	query := `INSERT INTO users(user_name, email, password, role, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		fmt.Println("Failed at query preperation stage: ", err)
		return errors.New("query failed at prepare stage")
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.UserName, u.Email, u.Password, u.Role, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		fmt.Println("Failed at query execution stage", err)
		return errors.New("query failed at execution stage")
	}
	return err
}

func (c Credentials) GetUserByUserName() (Credentials, error) {
	query := `SELECT id, user_name, password, role FROM users WHERE user_name = ?`
	row := database.DB.QueryRow(query, c.UserName)
	err := row.Scan(&c.ID, &c.UserName, &c.Password, &c.Role)
	if err != nil {
		return Credentials{}, err
	}
	return c, nil
}
