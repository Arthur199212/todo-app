package repository

import "github.com/jmoiron/sqlx"

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser() {}

func (r *AuthPostgres) GetUser() {}
