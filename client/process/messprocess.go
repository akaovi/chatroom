package process

import (
	"chatroom_sever/util"
	"encoding/json"
	"fmt"
	"net"
)

type Messprocess struct {
	Conn net.Conn
	Uid  int
}

func (m *Messprocess) GetallOnlineuser(mess util.Message) (err error) {

	users := make([]int, 0)

	err = json.Unmarshal([]byte(mess.Data), &users)
	if err != nil {
		return
	}

	fmt.Println("-----------------------------------------------------")
	for _, v := range users {
		fmt.Println("在线用户：", v)
	}
	fmt.Println("-----------------------------------------------------")

	return
}

func (m *Messprocess) AcceptMess(mess util.Message) (err error) {
	var asm util.AcceptMess

	err = json.Unmarshal([]byte(mess.Data), &asm)
	if err != nil {
		return
	}

	fmt.Println("-----------------------------------------------------")
	fmt.Println("用户为:", asm.From_Uid, "给你发消息了")
	fmt.Println("消息: ", asm.Data)
	fmt.Println("-----------------------------------------------------")
	return
}

func (m *Messprocess) LogOut() {
	fmt.Println("离线消息处理完成...")

}

func (m *Messprocess) GetSendMess(Uid int) (err error) {
	//
	return
}

func (m *Messprocess) AddFriend(mess util.Message) (err error) {
	// 判断是否添加好友

	var acmess util.AcceptMess

	err = json.Unmarshal([]byte(mess.Data), &acmess)
	if err != nil {
		return
	}

	fmt.Println("-----------------------------------------------------")
	fmt.Println("用户为:", acmess.From_Uid, "请求加你为好友")
	fmt.Println("验证消息为:", acmess.Data)
	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("请输入1或者2表示同意或者拒绝添加对方为好友")
	var keyvalue int = 1 // 不能解决在for下有开一个scanln
	// fmt.Scanln(&keyvalue)

	if keyvalue == 1 {
		// 同意添加对方为好友

		befriend := util.AddFriendMess{
			AdderUid: acmess.From_Uid,
			AddUid:   m.Uid,
		}

		data, err := json.Marshal(befriend)
		if err != nil {
			fmt.Println("添加失败，服务器报错")
			return err
		}

		remess := &util.Message{
			Type: util.AccpetBeFriendType,
			Data: string(data),
		}

		data, err = json.Marshal(remess)
		if err != nil {
			fmt.Println("添加失败，服务器报错")
			return err
		}

		tf := &util.Transfer{
			Conn: m.Conn,
		}

		tf.WritePkg(data)

	}

	return

}
