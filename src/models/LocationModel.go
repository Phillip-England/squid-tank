package models

import (
	"cfa-tools-api/src/e"
	"strconv"
	"strings"
)

type LocationModel struct {
	Id string `bson:"_id" json:"-"`
	User string `bson:"user" json:"user"`
	Name string `bson:"name" json:"name"`
	Number string `bson:"number" json:"number"`
}

func (v *LocationModel) Format() {
	v.Name = strings.ToLower(v.Name)
}

func (v *LocationModel) Validate() *e.HttpError {
	err := v.ValidateName()
	if err != nil {
		return err
	}
	err = v.ValidateNumber()
	if err != nil {
		return err
	}
	return nil
}

func (v *LocationModel) ValidateName() *e.HttpError {
	if len(v.Name) > 64 {
		return e.NewHttpError("name too long", 400)
	}
	if len(v.Name) < 5 {
		return e.NewHttpError("name too short", 400)
	}
	return nil
}

func (v *LocationModel) ValidateNumber() *e.HttpError {
	_, err := strconv.Atoi(v.Number)
	if err != nil {
		return e.NewHttpError("must provide a valid number", 400)
	}
	if len(v.Number) > 32 {
		return e.NewHttpError("number too long", 400)
	}
	if len(v.Number) < 5 {
		return e.NewHttpError("number too short", 400)
	}
	return nil
}