package todo

import (
	"regexp"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Id       string `json:"-" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password_hash"`
}

func (u User) Validate() error {
	return v.ValidateStruct(&u,
		v.Field(&u.Email, v.Required, is.Email),
		v.Field(&u.Password, v.Required, v.Length(6, 30)),
		v.Field(&u.Password, v.Required, v.Match(regexp.MustCompile("[A-Z]{1}")).Error("should have at least 1 upper case letter")),
		v.Field(&u.Password, v.Required, v.Match(regexp.MustCompile("[0-9]{1}")).Error("should have at least 1 number")),
		v.Field(&u.Password, v.Required, v.Match(regexp.MustCompile("[#?!@$%^&*-]{1}")).Error("should have at least 1 special character")),
	)
}
