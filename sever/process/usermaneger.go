package process

import (
	"fmt"
)

var (
	usermaneger *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	usermaneger = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

func (u *UserMgr) Addonlineuser(up *UserProcess) {
	u.onlineUsers[up.Uid] = up
}

func (u *UserMgr) Delonlineuser(uid int) {
	delete(u.onlineUsers, uid)
}

func (u *UserMgr) Getall_onlineuser() map[int]*UserProcess {
	return u.onlineUsers
}

func (u *UserMgr) GetuserByuid(uid int) (up *UserProcess, err error) {
	up, ok := u.onlineUsers[uid]

	if !ok {
		err = fmt.Errorf("该用户不在线")
		return
	}
	return

}

func (u *UserMgr) GetConn(uid int) *UserProcess {
	return u.onlineUsers[uid]
}
