package repository

import (
	"fmt"
	"todo-app/models"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf(`insert into %s (email, password_hash)
												values($1, $2) returning id`, usersTable)
	row := r.db.QueryRow(query, user.Email, user.Password)
	err := row.Scan(&id)
	return id, err
}

func (r *AuthPostgres) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("select * from %s where email=$1", usersTable)
	err := r.db.Get(&user, query, email)
	return user, err
}

func (r *AuthPostgres) GetUserById(id int) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("select * from %s where id=$1", usersTable)
	err := r.db.Get(&user, query, id)
	return user, err
}
