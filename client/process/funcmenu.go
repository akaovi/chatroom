package process

import (
	"chatroom_sever/util"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Menu struct {
	Conn net.Conn
	Uid  int
}

func (mn *Menu) Showfunc() {

	loop := true

	for loop {
		fmt.Println("-----------------------------------------------------")
		fmt.Println("-		1.显示在线好友列表  		-")
		fmt.Println("-		2.指定账号发送消息  		-")
		fmt.Println("-		3.添加好友          		-")
		fmt.Println("-		4.在线好友          		-")
		fmt.Println("-		5.退出登录          		-")
		fmt.Println("-		6.退出登录并退出系统		-")
		fmt.Println("-----------------------------------------------------")

		var key int
		fmt.Println("请输入(1~5): ")
		fmt.Scanln(&key)

		switch key {
		case 1:

			err := mn.client_getonlineuser()
			if err != nil {
				fmt.Println("err: ", err)
			}

		case 2:

			err := mn.sendmess()
			if err != nil {
				fmt.Println("消息发送失败 err: ", err)
			}

		case 3:

			err := mn.AddFriend()
			if err != nil {
				fmt.Println(err)
			}

		case 4:
			// 显示在线好友

		case 5:

			mn.logout()
			loop = false
			fmt.Printf("用户: %d 退出登录...\n", mn.Uid)

		case 6:

			mn.logout()
			fmt.Printf("用户: %d 退出登录并退出系统...\n", mn.Uid)
			os.Exit(0)

		default:
			fmt.Println("非法输入...")
		}
	}
}

func (mn *Menu) client_getonlineuser() (err error) {

	tf := &util.Transfer{
		Conn: mn.Conn,
	}

	mess := &util.Message{
		Type: util.MessOnlineType,
	}

	data, err := json.Marshal(mess)
	if err != nil {
		return
	}

	tf.WritePkg(data)
	return
}

func (mn *Menu) sendmess() (err error) {
	fmt.Println("请输入发送对象的账号: ")
	var to_uid int
	fmt.Scanln(&to_uid)

	fmt.Println("请输入发送内容(不支持换行): ")
	var data_string string
	fmt.Scanln(&data_string)

	// 写包
	sm := &util.SendMess{
		From_Uid: mn.Uid,
		To_Uid:   to_uid,
		Data:     data_string,
	}

	data, err := json.Marshal(sm)
	if err != nil {
		return
	}

	mess := &util.Message{
		Type: util.SendMessType,
		Data: string(data),
	}

	data, err = json.Marshal(mess)

	tf := &util.Transfer{
		Conn: mn.Conn,
	}

	tf.WritePkg(data)

	return
}

func (mn *Menu) logout() {
	lo := &util.Logout{
		Uid: mn.Uid,
	}

	data, err := Offlp.Sendprocess()
	if err != nil {
		fmt.Println(err)
	}

	lo.Data = string(data)

	data, err = json.Marshal(lo)
	if err != nil {
		fmt.Println(err)
	}

	mess := util.Message{
		Type: util.LogOutType,
		Data: string(data),
	}

	data, err = json.Marshal(mess)
	if err != nil {
		fmt.Println(err)
	}

	tf := &util.Transfer{
		Conn: mn.Conn,
	}

	tf.WritePkg(data)

	// 等待服务器做出 退出回应

	// mess, err = tf.ReadPkg()
	// if err != nil {
	// 	fmt.Println("离线消息处理失败...")
	// }

	// var relogout util.ReLogout

	// err = json.Unmarshal([]byte(mess.Data), &relogout)
	// if err != nil || relogout.Type != "OK" {
	// 	fmt.Println("离线消息处理失败...")
	// } else {
	// 	fmt.Println("离线消息处理正常完成...")
	// }

}

func (mn *Menu) AddFriend() (err error) {
	// 添加好友

	fmt.Println("请输入添加为好友的账号:")
	var frienduid int
	fmt.Scanln(&frienduid)

	fmt.Println("请输入你想对他(她)说的话:")
	var datastr string
	fmt.Scanln(&datastr)

	messfrienf := &util.AddFriendMess{
		AdderUid: mn.Uid,
		AddUid:   frienduid,
		Data:     datastr,
	}

	data, err := json.Marshal(messfrienf)
	if err != nil {
		fmt.Println("添加好友失败 err", err)
		return
	}

	mess := &util.Message{
		Type: util.AddFriendType,
		Data: string(data),
	}

	data, err = json.Marshal(mess)
	if err != nil {
		fmt.Println("添加好友失败...")
	}

	tf := &util.Transfer{
		Conn: mn.Conn,
	}

	tf.WritePkg(data)

	return

}
