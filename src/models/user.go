package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

//User describes a valid Nego user.
type User struct {
	*gorm.Model
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `validate:"required" json:"email" gorm:"index:user_email_index;unique;not_null"`
	Password string `validate:"required" json:"password"`
	Token    string `json:"token" gorm:"-"`
}

//NewUser returns a new instance of User.
func NewUser(name, lastName, password, email, token string) *User {
	return &User{
		Name:     name,
		LastName: lastName,
		Password: password,
		Email:    email,
		Token:    token,
	}
}

//Valid checks if the model has errors.
func (u *User) Valid() error {
	v := validator.New()
	return v.Struct(u)
}
