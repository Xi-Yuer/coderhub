package validator

import (
	"errors"
	"strings"
)

type Validator struct {
	errors []string
}

func New() *Validator {
	return &Validator{errors: make([]string, 0)}
}

func (v *Validator) Clear() {
	v.errors = make([]string, 0)
}

func (v *Validator) Username(username string) *Validator {
	if !UsernameRegex.MatchString(username) {
		v.errors = append(v.errors, "用户名格式错误：需要3-32位字母、数字或下划线")
	}
	return v
}

func (v *Validator) Password(password string) *Validator {
	if !PasswordRegex.MatchString(password) {
		v.errors = append(v.errors, "密码格式错误：需要6-32位字母或数字")
	}
	return v
}

func (v *Validator) Check() error {
	if len(v.errors) > 0 {
		defer v.Clear()
		return errors.New(strings.Join(v.errors, "; "))
	}
	return nil
}
