package repository

import (
	"fmt"
	"todo-app"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf(`insert into %s (email, password_hash)
												values($1, $2) returning id`, usersTable)
	row := r.db.QueryRow(query, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUserByEmail(email string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("select * from %s where email=$1", usersTable)
	err := r.db.Get(&user, query, email)
	return user, err
}
