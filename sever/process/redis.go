package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/garyburd/redigo/redis"
)

var Pool *redis.Pool

func init() {
	Pool = &redis.Pool{
		MaxIdle:     7,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "10.81.35.111:6379")
		},
	}
}

func New_UserId() uint32 {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(80000000)
	return uint32(a)
}

func OldUser_SignIn(uid int, UserPwd string) (code int, UserNickName string, err error) {
	conn := Pool.Get()
	defer conn.Close()

	v1, err := conn.Do("exists", uid)
	if err != nil {
		code = 500
		return
	}
	if int(v1.(int64)) != 1 {
		code = 300
		err = nil
		return
	}
	v2, err := redis.String(conn.Do("hget", uid, "pwd"))
	if err != nil {
		code = 500
		return
	}
	if v2 != UserPwd {
		code = 400
		err = nil
		return
	}

	v3, err := redis.String(conn.Do("hget", uid, "nickname"))
	if err != nil {
		code = 500
		return
	}
	UserNickName = v3

	code = 200
	err = nil

	return

}

func New_User_Create(nickname string, pwd string) (code int, uid int, err error) {
	conn := Pool.Get()
	defer conn.Close()

	var usid uint32
	loop1 := true
	for loop1 {
		usid = New_UserId()
		v, err := conn.Do("exists", usid)
		if err != nil {
			code = 500
			break
		} else {
			if int(v.(int64)) == 0 {
				loop1 = false
				uid = int(usid)
			}
		}
	}

	_, err = conn.Do("hmset", uid, "nickname", nickname, "pwd", pwd)
	if err != nil {
		fmt.Println("注册失败... err: ", err)
		code = 500
		return
	} else {
		_, err = conn.Do("save")
		if err != nil {
			fmt.Println("注册失败 save err: ", err)
			code = 500
			return
		}
	}

	code = 200

	return

}

func Friends(uid1, uid2 int) (err error) {
	conn := Pool.Get()
	defer conn.Close()

	v2, err := redis.String(conn.Do("hget", uid1, "friends"))
	if err != nil {
		oldf := []int{uid2}
		data, err := json.Marshal(oldf)
		if err != nil {
			return err
		}

		_, err = conn.Do("hset", uid1, "friends", string(data))
		if err != nil {
			return err
		}

		_, err = conn.Do("save")
		if err != nil {
			err = errors.New("添加好友失败")
			return err
		}
	} else {
		var oldf []int
		err = json.Unmarshal([]byte(v2), &oldf)
		if err != nil {
			return
		}

		oldf = append(oldf, uid2)

		data, err := json.Marshal(oldf)
		if err != nil {
			return err
		}

		_, err = conn.Do("hset", uid1, "friends", string(data))
		if err != nil {
			return err
		}

		_, err = conn.Do("save")
		if err != nil {
			err = errors.New("添加好友失败")
			return err
		}

	}

	return
}
