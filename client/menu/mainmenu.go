package menu

import (
	"chatroom_sever/client/process"
	"fmt"
)

type Chatroom struct {
	key  string
	loop bool
}

func NewChatroom() *Chatroom {
	return &Chatroom{
		key:  "",
		loop: true,
	}
}

func (f *Chatroom) exit() {
	justice := ""
	fmt.Println("真的要退出吗 y/n")
	for {
		fmt.Scanln(&justice)
		if justice == "y" || justice == "Y" {
			f.loop = false
			break
		} else if justice == "n" || justice == "N" {
			fmt.Println("感谢您的不退出,老泪纵横！")
			break
		} else {
			fmt.Println("输入有误，请重新输入！")
		}
	}
}

func (f *Chatroom) MainMenu() {

	up := &process.Userprocess{}

	for {
		fmt.Println("----------------欢迎来到飞来飞去聊天室----------------")
		fmt.Println("                        菜  单")
		fmt.Println("             	      1.登  陆")
		fmt.Println("                      2.注  册")
		fmt.Println("                      3.退  出")
		fmt.Println("------------------------------------------------------")

		fmt.Println("请输入(1~3):")
		fmt.Scanln(&f.key)

		switch f.key {
		case "1":
			up.Sign_In()
		case "2":
			up.Sign_Up()
		case "3":
			f.exit()
		default:
			fmt.Println("输入有误，请输入(1~3):")
		}
		if !f.loop {
			fmt.Println("退出飞来飞去聊天室！")
			break
		}
	}
}
