package util

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (t *Transfer) WritePkg(data []byte) (err error) {
	// 消息长度content length
	pkgLen := uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(t.Buf[0:4], pkgLen)
	n, err := t.Conn.Write(t.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("消息传输中断(length err)... err: ", err)
		return
	}

	// 消息本身
	n, err = t.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("Write data err: ", err)
		return
	}
	return
}

func (t *Transfer) ReadPkg() (mess Message, err error) {

	// buf := make([]byte, 8096)
	// fmt.Println("读取客户端: ", t.Conn.LocalAddr().String(), "数据中...")
	_, err = t.Conn.Read(t.Buf[:4])
	if err != nil {
		fmt.Println("Read err: ", err)
		return
	}

	// var pkgLen uint32
	pkgLen := binary.BigEndian.Uint32(t.Buf[0:4])
	n, err := t.Conn.Read(t.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}

	err = json.Unmarshal(t.Buf[:pkgLen], &mess)
	if err != nil {
		fmt.Println("Unmarshal err: ", err)
		return
	}
	return

}
