package repository

import (
	"database/sql"
)

type Repository interface {
	GetUserNameByID(id int64) (string, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserNameByID(id int64) (string, error) {
	var username string
	err := r.db.QueryRow("SELECT username FROM users WHERE id = ?", id).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}
