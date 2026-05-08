package domain

import (
	"github.com/aliworkshop/errors"
	"github.com/aliworkshop/gateway/v2"
	"github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Username string
	Password string
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"required"`
}

func (r RefreshRequest) Validate(v *validator.Validate, _ gateway.Language) errors.ErrorModel {
	if err := v.Struct(r); err != nil {
		return errors.Validation(err)
	}
	return nil
}

type PostRequest struct {
	Id     string `json:"id" param:"id"`
	Name   string `json:"name" form:"name"`
	Data   string `json:"data"`
	Value  string `json:"value"`
	Number uint64 `json:"number"`
}

func (PostRequest) Validate(*validator.Validate, gateway.Language) errors.ErrorModel {
	return nil
}