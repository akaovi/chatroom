package process

import (
	"chatroom_sever/util"
	"fmt"
	"net"
)

type SeverProcess struct {
	Conn net.Conn
}

func (s SeverProcess) Severprocess(mess util.Message) {

	switch mess.Type {

	case util.MessLoginType:
		up := &UserProcess{
			Conn: s.Conn,
		}
		err := up.Login_sever(mess)
		if err != nil {
			fmt.Println(err)
		}

	case util.MessRegisterType:
		up := &UserProcess{
			Conn: s.Conn,
		}
		err := up.Register(mess)
		if err != nil {
			fmt.Println(err)
		}

	case util.MessOnlineType:
		up := &UserProcess{
			Conn: s.Conn,
		}
		err := up.GetOnlineUser()
		if err != nil {
			fmt.Println(err)
		}

	case util.SendMessType:
		up := &UserProcess{
			Conn: s.Conn,
		}
		err := up.SendMessage(mess)
		if err != nil {
			fmt.Println(err)
		}

	case util.LogOutType:
		up := &UserProcess{
			Conn: s.Conn,
		}
		err := up.Userlogout(mess)
		if err != nil {
			fmt.Println(err)
		}

	case util.AddFriendType:
		up := &UserProcess{
			Conn: s.Conn,
		}
		err := up.Addfriend(mess)
		if err != nil {
			fmt.Println(err)
		}
	case util.AccpetBeFriendType:
		up := &UserProcess{
			Conn: s.Conn,
		}
		err := up.BeFriend(mess)
		if err != nil {
			return
		}

	default:
		fmt.Println("无法识别客户端发送的消息类型...")
	}
}
