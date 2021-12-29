package http

import (
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
	"net/http"
)

func configPlayerRoutes() {
	http.HandleFunc("/player", func(w http.ResponseWriter, r *http.Request) {
		setupHeader(w, r)
		if r.Method == http.MethodGet {
			RenderDataJson(w, g.Config().Player)
		}
	})
}
