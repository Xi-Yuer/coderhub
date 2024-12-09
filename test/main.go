package main

import (
	"coderhub/shared/GoMail"
)

func main() {
	g := GoMail.NewGoMail()
	g.SendWithHTML("2214380963@qq.com", "邮箱密码重置", "https://www.baidu.com")
}
