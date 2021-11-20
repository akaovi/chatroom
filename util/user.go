package util

type User struct {
	UserId       int    `json:"userId"`
	UserNickName string `json:"userNickName"`
	UserPwd      string `json:"userPwd"`
}

func New_User() *User {
	return &User{
		UserId:       0,
		UserNickName: "",
		UserPwd:      "",
	}
}
