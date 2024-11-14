package Validator

import "regexp"

var (
	UsernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,32}$`)
	PasswordRegex = regexp.MustCompile(`^[A-Za-z\d]{6,32}$`)
	EmailRegex    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)
