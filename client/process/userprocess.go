package process

import (
	"chatroom_sever/util"
	"encoding/json"
	"fmt"
	"net"
)

type Userprocess struct {
	//
}

// func KeepAlive(conn net.Conn) {
// 	tf := &util.Transfer{
// 		Conn: conn,
// 	}
// 	for {
// 		tf.ReadPkg()
// 	}
// }

func (u *Userprocess) Sign_In() {

	messlogin := &util.MessLogin{}

	loop := true
	for loop {
		fmt.Println("请输入你的账号：")
		_, err := fmt.Scanln(&messlogin.Uid)
		if err != nil {
			fmt.Println("输入有误...")
		} else {
			loop = false
		}
	}
	fmt.Println("请输入你的密码：")
	fmt.Scanln(&messlogin.Pwd)

	// 连接服务器

	conn, err := net.Dial("tcp", "10.81.35.111:9977")
	if err != nil {
		fmt.Println("Dial err: ", err)
		return
	}
	defer conn.Close()

	var mess util.Message
	mess.Type = util.MessLoginType

	data, err := json.Marshal(messlogin)
	if err != nil {
		fmt.Println("json.Marshal(User) err: ", err)
		return
	}
	mess.Data = string(data)

	data, err = json.Marshal(mess)
	if err != nil {
		fmt.Println("登陆失败... Marshal err: ", err)
		return
	}

	transfer := &util.Transfer{
		Conn: conn,
	}

	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg err: ", err)
		return
	}

	// 接收服务器的返回

	var Remess util.ReMessLogin

	mess, err = transfer.ReadPkg()
	if err != nil {
		fmt.Println("ReadPkg err: ", err)
		return
	}

	err = json.Unmarshal([]byte(mess.Data), &Remess)
	if err != nil {
		fmt.Println("Unmarshal err: ", err)
		return
	}

	if Remess.Code == 200 {
		fmt.Println("登陆成功...")
		fmt.Println("欢迎回来: ", Remess.Nickname)

		// 保持与服务器通信
		go func() {
			sp := &ClientProcess{
				Conn: conn,
				Uid:  messlogin.Uid,
			}
			for {
				mess, err := transfer.ReadPkg()
				if err != nil {
					fmt.Println("与服务器断开连接 err: ", err)
					return
				}

				sp.Clientprocess(mess)
			}
		}()

		// 登陆后的菜单
		menu := &Menu{
			Conn: conn,
			Uid:  messlogin.Uid,
		}
		menu.Showfunc()

	} else if Remess.Code == 300 {
		fmt.Println("该账号未注册...")
	} else if Remess.Code == 400 {
		fmt.Println("账号或密码错误...")
	} else {
		fmt.Println("服务器维护中...")
	}

}

func (u *Userprocess) Sign_Up() {

	mg := &util.MessRegister{}

	temp := ""
	loop1 := true
	loop2 := true
	for loop1 {
		fmt.Println("请输入你的昵称：")
		_, err := fmt.Scanln(&mg.Nickname)
		if err != nil {
			fmt.Println("输入有误...")
		} else {
			loop1 = false
		}
	}
	for loop2 {
		fmt.Println("请输入你新密码：")
		fmt.Scanln(&temp)
		fmt.Println("请再次输入你的密码：")
		fmt.Scanln(&mg.Pwd)

		if mg.Pwd == temp {
			fmt.Println("昵称：", mg.Nickname, "密码：", mg.Pwd)
			loop2 = false
		} else {
			fmt.Printf("两次密码输入不一致，请重新输入...\n\n")
		}
	}

	// 连接服务器

	conn, err := net.Dial("tcp", "10.81.35.111:9977")
	if err != nil {
		fmt.Println("Dial err: ", err)
		return
	}
	defer conn.Close()

	var mess util.Message
	mess.Type = util.MessRegisterType

	data, err := json.Marshal(mg)
	if err != nil {
		fmt.Println("json.Marshal(mg) err: ", err)
		return
	}
	mess.Data = string(data)

	data, err = json.Marshal(mess)
	if err != nil {
		fmt.Println("注册失败... Marshal err: ", err)
		return
	}

	ts := &util.Transfer{
		Conn: conn,
	}

	err = ts.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg err: ", err)
		return
	}

	// 接收服务器的消息

	remess := &util.ReMessRegister{}

	mess, err = ts.ReadPkg()
	if err != nil {
		fmt.Println("ReadPkg err: ", err)
		return
	}

	err = json.Unmarshal([]byte(mess.Data), &remess)
	if err != nil {
		fmt.Println("Unmarshal err: ", err)
		return
	}

	if remess.Code == 200 {
		fmt.Println("注册成功...")
		fmt.Println("请记住你的账号: ", remess.Uid)
		fmt.Println("因为凭账号登录哦, 小可爱, ...")

	} else {
		fmt.Println("注册失败， 服务器错误...")
	}

}
