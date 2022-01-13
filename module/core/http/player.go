package http

import (
	"encoding/json"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/core/g/db"
	bolt "go.etcd.io/bbolt"
	"net/http"
)

func configPlayerRoutes() {
	http.HandleFunc("/player", func(w http.ResponseWriter, r *http.Request) {
		setupHeader(w, r)
		rtu := make([]model.Player, 0)
		if r.Method == http.MethodGet {
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
		if r.Method == http.MethodPost {

		}
	})
}
