package main

import (
	"coderhub/shared/bcryptUtil"
	"fmt"
)

func main() {
	passwordHash, _ := bcryptUtil.PasswordHash("123")
	err := bcryptUtil.CompareHashAndPassword("123", passwordHash)
	if err {
		fmt.Println(err)
		return
	}
	fmt.Println("ok")
}
