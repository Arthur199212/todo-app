package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Id       int    `json:"-" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password_hash"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(
			&u.Password,
			validation.Required,
			validation.Length(6, 30),
			validation.Match(regexp.MustCompile("[A-Z]{1}")).Error("should have at least 1 upper case letter"),
			validation.Match(regexp.MustCompile("[0-9]{1}")).Error("should have at least 1 number"),
			validation.Match(regexp.MustCompile("[#?!@$%^&*-]{1}")).Error("should have at least 1 special character"),
		),
	)
}
