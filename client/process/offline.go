package process

import (
	"encoding/json"
	"errors"
)

var (
	Offlp *OfflineMessProcess
)

type OfflineMessProcess struct {
	OfflineMess map[string][]string
}

func init() {
	Offlp = &OfflineMessProcess{
		OfflineMess: make(map[string][]string, 1024),
	}
}

func (op *OfflineMessProcess) Sendprocess() (data []byte, err error) {
	// 在用户离线时将待处理消息类型推送给服务器 暂时存储 等用户上线时 在推送给用户

	// 打包Offline 给Logout
	if len(op.OfflineMess) == 0 {
		data = []byte{}
		err = nil
		return
	}

	data, err = json.Marshal(op.OfflineMess)
	if err != nil {
		err = errors.New("离线信息保存失败")
		return
	}

	return

}
