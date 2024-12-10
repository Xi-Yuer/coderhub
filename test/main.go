package main

import (
	"coderhub/shared/messaging"
)

func main() {
	g := messaging.NewGoMail()
	err := g.SendWithHTML("2214380963@qq.com", "邮箱密码重置", "https://www.baidu.com")
	if err != nil {
		return
	}
}
