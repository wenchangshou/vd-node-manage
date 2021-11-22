package g

import (
	"encoding/json"
	"github.com/denisbrodbeck/machineid"
	"github.com/toolkits/file"
	"github.com/wenchangshou2/vd-node-manage/common/safety"
	"log"
	"sync"
)

type HardwareConfig struct {
	ID string `json:"id"`
}

var (
	HardwareFile string
	hardware     *HardwareConfig
	hardwareLock = new(sync.RWMutex)
	key          = []byte("123456781234567812345678")
)

func Hardware() *HardwareConfig {
	hardwareLock.RLock()
	defer hardwareLock.RUnlock()
	return hardware
}
func ParseHardware(cfg string) {
	if cfg == "" {
		log.Fatalln("use -d to specify hardware file")
	}
	if !file.IsExist(cfg) {
		id, err := machineid.ProtectedID("gateway")
		if err != nil {
			log.Fatalln("get machine id fail")
		}
		h := HardwareConfig{ID: id}
		b, _ := json.Marshal(h)
		c := safety.Encrypt(key, string(b))
		file.WriteBytes(cfg, []byte(c))
	}
	HardwareFile = cfg
	hardwareContent, err := file.ToString(cfg)
	hardwareContent= safety.Decrypt(key, string(hardwareContent))

	if err != nil {
		log.Fatalln("read hardware file:", cfg, "fail:", err)
	}
	var c HardwareConfig
	err = json.Unmarshal([]byte(hardwareContent), &c)
	if err != nil {
		log.Fatalln("parse hardware file", cfg, "error:", err.Error())
	}
	hardwareLock.Lock()
	defer hardwareLock.Unlock()
	hardware = &c
	log.Println("g.ParseHardware ok,file", cfg)
}
