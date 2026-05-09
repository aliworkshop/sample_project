package app

import (
	"reflect"
	"regexp"
	"strings"
	"sync"
)

type validate struct {
	once      sync.Once
	validator *validator.Validate
}

func (a *App) NewValidator() *validate {
	v := new(validate)
	v.once.Do(func() {
		v.validator = a.engine.(interface {
			Validator() *validator.Validate
		}).Validator()
		v.validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		err := v.validator.RegisterValidation("password_validator", a.ValidatePassword)
		a.panicOnErr(err)
	})
	a.validator = v.validator
	return v
}

func (a *App) ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	const (
		minPasswordLength = 8
		maxPasswordLength = 64
	)
	if len(password) < minPasswordLength || len(password) > maxPasswordLength {
		return false
	}
	hasLower := regexp.MustCompile(`[a-z]`).MatchString
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString
	hasDigit := regexp.MustCompile(`\d`).MatchString
	hasSpecial := regexp.MustCompile(`[@$!%*?&#]`).MatchString
	allowedChars := regexp.MustCompile(`^[A-Za-z\d@$!%*?&#]+$`).MatchString
	if !hasLower(password) || !hasUpper(password) || !hasDigit(password) || !hasSpecial(password) || !allowedChars(password) {
		return false
	}
	return true
}
