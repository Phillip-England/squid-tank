package models

import (
	"cfa-tools-api/src/e"
	"cfa-tools-api/src/lib"
	"net/mail"
	"strings"
)

type UserModel struct {
	Email string `bson:"email"`
	Password string `bson:"password"`
}

func (v UserModel) Format() {
	v.Email = strings.ToLower(v.Email)
}

func (v UserModel) Validate() *e.HttpError {
	err := v.ValidateEmail()
	if err != nil {
		return err
	}
	err = v.ValidatePassword()
	if err != nil {
		return err
	}
	return nil
}

func (v UserModel) ValidateEmail() *e.HttpError {
	_, err := mail.ParseAddress(v.Email)
	if err != nil {
		return e.NewHttpError("invalid email", 400)
	}
	return nil
}

func (v UserModel) ValidatePassword() *e.HttpError {
	if len(v.Password) > 64 {
		return e.NewHttpError("password too long", 400)
	}
	if len(v.Password) < 5 {
		return e.NewHttpError("password too short", 400)
	}
	return nil
}

func (v UserModel) HashPassword() *e.HttpError {
	hashed_password, err := lib.HashString(v.Password)
	if err != nil {
		return e.NewHttpError("internal server error", 500)
	}
	v.Password = hashed_password
	return nil
}