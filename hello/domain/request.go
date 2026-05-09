package domain

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type LoginRequest struct {
	Username string
	Password string
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"required"`
}

func (r RefreshRequest) Validate(v *validator.Validate, lang gateway.Language) errors.ErrorModel {
	if err := v.Struct(r); err != nil {
		e := errors.Validation()
		for _, ev := range err.(validator.ValidationErrors) {
			msg, _ := lang.Localize(&i18n.LocalizeConfig{
				MessageID:      ev.Tag(),
				TemplateData:   map[string]any{"field": ev.Field(), "param": ev.Param()},
				DefaultMessage: &i18n.Message{ID: ev.Tag(), Other: ev.Error()},
			})
			e = e.WithProperty(ev.Field(), msg)
		}
		return e
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

func (r PostRequest) Validate(v *validator.Validate, lang gateway.Language) errors.ErrorModel {
	if err := v.Struct(r); err != nil {
		e := errors.Validation()
		for _, ev := range err.(validator.ValidationErrors) {
			msg, _ := lang.Localize(&i18n.LocalizeConfig{
				MessageID:      ev.Tag(),
				TemplateData:   map[string]any{"field": ev.Field(), "param": ev.Param()},
				DefaultMessage: &i18n.Message{ID: ev.Tag(), Other: ev.Error()},
			})
			e = e.WithProperty(ev.Field(), msg)
		}
		return e
	}
	return nil
}
