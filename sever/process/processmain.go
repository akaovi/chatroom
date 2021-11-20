package process

import (
	"chatroom_sever/util"
	"fmt"
	"net"
)

type ProcessMain struct {
	Conn net.Conn
}

func (p ProcessMain) Processmain() {

	sp := &SeverProcess{
		Conn: p.Conn,
	}

	ts := &util.Transfer{
		Conn: p.Conn,
	}

	for {
		mess, err := ts.ReadPkg()
		if err != nil {
			fmt.Println("repkg err: ", err)
			return
		}

		fmt.Println(mess)

		sp.Severprocess(mess)

	}
}
