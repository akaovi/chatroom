package process

type OfflineProcess struct {
	OfflineUser map[int]map[string][]string
}

var (
	Offu *OfflineProcess
)

func init() {
	Offu = &OfflineProcess{
		OfflineUser: make(map[int]map[string][]string, 1024),
	}
}
