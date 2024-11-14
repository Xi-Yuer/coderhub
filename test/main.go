package main

import (
	"coderhub/shared/BcryptUtil"
	"fmt"
)

func main() {
	passwordHash, _ := BcryptUtil.PasswordHash("123")
	err := BcryptUtil.CompareHashAndPassword("123", passwordHash)
	if err {
		fmt.Println(err)
		return
	}
	fmt.Println("ok")
}
