package http

import (
	"encoding/json"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/core/g/db"
	bolt "go.etcd.io/bbolt"
	"net/http"
)

func GetPlayer(w http.ResponseWriter) {
	rtu := make([]model.Player, 0)
	err := db.GDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("player"))
		if bucket == nil {
			return nil
		}
		err := bucket.ForEach(func(k, v []byte) error {
			p := model.Player{}
			err := json.Unmarshal(v, &p)
			if err != nil {
				return err
			}
			rtu = append(rtu, p)
			return nil
		})
		return err
	})
	if err != nil {
		RenderMsgJson(w, "获取数据失败", err)
		return
	}
	RenderDataJson(w, rtu)
	return
}
func AddPlayer(w http.ResponseWriter, r *http.Request) {

}
func configPlayerRoutes() {
	http.HandleFunc("/player", func(w http.ResponseWriter, r *http.Request) {
		setupHeader(w, r)
		if r.Method == http.MethodGet {
			GetPlayer(w)
		}
		if r.Method == http.MethodPost {
			AddPlayer(w, r)
		}
	})
}
