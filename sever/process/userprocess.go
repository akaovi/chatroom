package process

import (
	"chatroom_sever/util"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	Uid  int
}

// 登录

func (u *UserProcess) Login_sever(mess util.Message) (err error) {

	var ml util.MessLogin
	err = json.Unmarshal([]byte(mess.Data), &ml)
	if err != nil {
		fmt.Println("登录失败... Unmarshal err: ", err)
		return
	}

	code, nickname, err := OldUser_SignIn(ml.Uid, ml.Pwd)
	if err != nil {
		fmt.Println("账号密码验证错误...")
	}

	if code == 200 {
		u.Uid = ml.Uid
		usermaneger.Addonlineuser(u)
	}

	rml := &util.ReMessLogin{
		Code:     code,
		Nickname: nickname,
	}

	data, err := json.Marshal(rml)
	if err != nil {
		return
	}

	remess := &util.Message{
		Type: util.ReMessLoginType,
		Data: string(data),
	}

	data, err = json.Marshal(remess)
	if err != nil {
		return
	}

	ts := util.Transfer{
		Conn: u.Conn,
	}

	err = ts.WritePkg(data)
	if err != nil {
		return
	}

	return
}

// 注册

func (u *UserProcess) Register(mess util.Message) (err error) {

	var mg util.MessRegister
	err = json.Unmarshal([]byte(mess.Data), &mg)
	if err != nil {
		fmt.Println("注册失败... Unmarshal err: ", err)
		return
	}

	code, uid, err := New_User_Create(mg.Nickname, mg.Pwd)
	if err != nil {
		fmt.Println(err)
	}

	rmg := &util.ReMessRegister{
		Code: code,
		Uid:  uid,
	}

	data, err := json.Marshal(rmg)
	if err != nil {
		return
	}

	remess := &util.Message{
		Type: util.ReMessRegisterType,
		Data: string(data),
	}

	data, err = json.Marshal(remess)
	if err != nil {
		return
	}

	ts := &util.Transfer{
		Conn: u.Conn,
	}

	err = ts.WritePkg(data)
	if err != nil {
		return
	}

	return
}

// 返回在线用户列表

func (u *UserProcess) GetOnlineUser() (err error) {

	users := make([]int, 0)
	for i := range usermaneger.onlineUsers {
		users = append(users, i)
	}

	data, err := json.Marshal(users)
	if err != nil {
		return
	}

	mess := &util.Message{
		Type: util.MessOnlineType,
		Data: string(data),
	}

	tf := util.Transfer{
		Conn: u.Conn,
	}

	data, err = json.Marshal(mess)
	if err != nil {
		return
	}

	tf.WritePkg(data)

	return
}

func (u *UserProcess) SendMessage(mess util.Message) (err error) {
	var sm util.SendMess
	err = json.Unmarshal([]byte(mess.Data), &sm)
	if err != nil {
		return
	}

	connstruct := usermaneger.GetConn(sm.To_Uid)
	connstruct.sendMessage(sm.From_Uid, sm.Data)

	return

}

func (u *UserProcess) sendMessage(fromuid int, datastring string) (err error) {
	am := &util.AcceptMess{
		From_Uid: fromuid,
		Data:     datastring,
	}

	data, err := json.Marshal(am)
	if err != nil {
		return
	}

	mess := &util.Message{
		Type: util.AcceptMessType,
		Data: string(data),
	}

	data, err = json.Marshal(mess)
	if err != nil {
		return
	}

	tf := &util.Transfer{
		Conn: u.Conn,
	}

	tf.WritePkg(data)

	return
}

func (u *UserProcess) Userlogout(mess util.Message) (err error) {
	// 退出登录
	var lo util.Logout

	err = json.Unmarshal([]byte(mess.Data), &lo)
	if err != nil {
		return
	}

	usermaneger.Delonlineuser(lo.Uid)

	if lo.Data != string([]byte{}) {
		var map_off map[string][]string

		err = json.Unmarshal([]byte(lo.Data), &map_off)
		if err != nil {
			fmt.Println("离线消息处理失败...")
		}

		Offu.OfflineUser[u.Uid] = map_off
	}

	remess := &util.Message{
		Type: util.ReLogOutType,
		Data: "OK",
	}

	data, err := json.Marshal(remess)

	tf := &util.Transfer{
		Conn: u.Conn,
	}

	tf.WritePkg(data)

	return
}

func (u *UserProcess) Addfriend(mess util.Message) (err error) {
	// 添加好友

	var addfriendmess util.AddFriendMess

	err = json.Unmarshal([]byte(mess.Data), &addfriendmess)
	if err != nil {
		return
	}

	add_userpj, err := usermaneger.GetuserByuid(addfriendmess.AddUid)
	if err != nil {
		Offu.OfflineUser[addfriendmess.AddUid][util.AddFriendType] = append(Offu.OfflineUser[addfriendmess.AddUid][util.AddFriendType], mess.Data)
	}

	tf := &util.Transfer{
		Conn: add_userpj.Conn,
	}

	acmess := &util.AcceptMess{
		From_Uid: addfriendmess.AdderUid,
		Data:     addfriendmess.Data,
	}

	data, err := json.Marshal(acmess)
	if err != nil {
		return
	}

	mess1 := util.Message{
		Type: util.AskBeFriendType,
		Data: string(data),
	}

	data, err = json.Marshal(mess1)
	if err != nil {
		return
	}

	tf.WritePkg(data)

	return

}

func (u *UserProcess) BeFriend(mess util.Message) (err error) {
	// 成为好友

	var adf util.AddFriendMess

	err = json.Unmarshal([]byte(mess.Data), &adf)
	if err != nil {
		return
	}

	err = Friends(adf.AddUid, adf.AdderUid)
	if err != nil {
		return
	}
	err = Friends(adf.AdderUid, adf.AddUid)
	if err != nil {
		return
	}

	return

}
