package process

import (
	"chatroom_sever/util"
	"fmt"
	"net"
)

type ClientProcess struct {
	Conn net.Conn
	Uid  int
}

func (s *ClientProcess) Clientprocess(mess util.Message) {

	switch mess.Type {

	case util.MessOnlineType:
		mp := &Messprocess{
			Conn: s.Conn,
		}
		err := mp.GetallOnlineuser(mess)
		if err != nil {
			fmt.Println(err)
		}

	case util.AcceptMessType:
		mp := &Messprocess{
			Conn: s.Conn,
		}
		err := mp.AcceptMess(mess)
		if err != nil {
			fmt.Println("接收消息失败...")
		}

	case util.ReLogOutType:
		mp := &Messprocess{
			Conn: s.Conn,
		}
		mp.LogOut()

	case util.AskBeFriendType:
		mp := &Messprocess{
			Conn: s.Conn,
			Uid:  s.Uid,
		}
		mp.AddFriend(mess)

	default:
		fmt.Println("无法识别客户端发送的消息类型...")
	}
}
