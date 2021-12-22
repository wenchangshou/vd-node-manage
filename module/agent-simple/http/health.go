package http

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/buff"
	"io/ioutil"
	"net/http"
)

func configHealthRoutes() {
	router := httprouter.Router{}
	router.GET("/health", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Write([]byte("ok"))
	})
	router.GET("/window/:wid/health", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		buff.GWindowGlobalStatus.SetWindowHealth(params.ByName("wid"))
	})
	router.POST("/window/:wid/report", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		s, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		buff.GWindowGlobalStatus.SetWindowStatus(ps.ByName("wid"), string(s))
	})

}
