package util

var (
	MessLoginType      = "MessLogin"
	MessRegisterType   = "MessRegister"
	ReMessLoginType    = "ReMessLogin"
	ReMessRegisterType = "ReMessRegister"
	LogOutType         = "LogOut"
	ReLogOutType       = "ReLogOut"
	MessOnlineType     = "MessOnline"
	SendMessType       = "SendMess"
	AcceptMessType     = "AcceptMess"
	AddFriendType      = "AddFriend"
	AskBeFriendType    = "AskBeFriend"
	AccpetBeFriendType = "AccpetBeFriend"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// 登录与登录返回

type MessLogin struct {
	Uid int    `json:"uid"`
	Pwd string `json:"pwd"`
}

type ReMessLogin struct {
	Code     int    `json:"code"` // 200 登录成功  300 未注册 400 账号密码错误 500 其他错误
	Nickname string `json:"nickname"`
}

// 注册与注册返回

type MessRegister struct {
	Nickname string `json:"nickname"`
	Pwd      string `json:"pwd"`
}

type ReMessRegister struct {
	Code int `json:"code"` // 200 注册成功  500 注册失败,其他错误
	Uid  int `json:"uid"`
}

type ReMessOnline struct {
	Data string `json:"data"`
}

type SendMess struct {
	From_Uid int    `json:"from_Uid"`
	To_Uid   int    `json:"to_Uid"`
	Data     string `json:"data"`
}

type AcceptMess struct {
	From_Uid int    `json:"from_Uid"`
	Data     string `json:"data"`
}

type ReSendMess struct {
	From_Uid int    `json:"from_Uid"`
	To_Uid   int    `json:"to_Uid"`
	Data     string `json:"data"`
}

type Logout struct {
	Uid  int    `json:"uid"`
	Data string `json:"data"`
}

type ReLogout struct {
	Type string `json:"type"`
}

type AddFriendMess struct {
	AdderUid int    `json:"adderUid"`
	AddUid   int    `json:"addUid"`
	Data     string `json:"data"` // 验证消息
}

// type AcceptFriendMess struct {
// 	AdderUid int    `json:"adderUid"`
// 	AddUid   int    `json:"addUid"`
// 	Data     string `json:"data"` // 验证消息
// }
