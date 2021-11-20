package process

type Friends struct {
	Onlinefriends map[int]string
	AllFriends    map[int]string
}

var (
	Friendslst *Friends
)

func init() {
	Friendslst = &Friends{
		Onlinefriends: make(map[int]string),
		AllFriends:    make(map[int]string),
	}
}

func (f *Friends) DelOnlineFriend(uid int) {
	delete(f.Onlinefriends, uid)
}

func (f *Friends) AddOnlinefriend(uid int, nickname string) {
	f.Onlinefriends[uid] = nickname
}
