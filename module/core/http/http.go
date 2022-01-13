package http

import (
	"encoding/json"
	"github.com/wenchangshou/vd-node-manage/module/core/g"
	"log"
	"net/http"
)

type Dto struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func init() {
	configHealthRoutes()
	configSystemRoutes()
	configPlayerRoutes()
}
func RenderJson(w http.ResponseWriter, v interface{}) {
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bs)
}
func RenderDataJson(w http.ResponseWriter, data interface{}) {
	RenderJson(w, Dto{Code: 0, Msg: "success", Data: data})
}
func RenderCustomMsgJson(w http.ResponseWriter, code int, msg string) {

	rtu := make(map[string]interface{})
	rtu["code"] = code
	rtu["msg"] = msg
	RenderJson(w, rtu)
}
func RenderMsgJson(w http.ResponseWriter, msg string, err error) {
	rtu := make(map[string]interface{})
	rtu["code"] = 0
	rtu["msg"] = msg
	if err != nil {
		rtu["code"] = 400
		rtu["msg"] = err.Error()
	}
	RenderJson(w, map[string]string{"msg": msg})
}
func AutoRender(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		RenderMsgJson(w, err.Error(), err)
		return
	}
	RenderDataJson(w, data)
}
func setupHeader(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
func Start() {
	if !g.Config().Http.Enabled {
		return
	}
	addr := g.Config().Http.Listen
	if addr == "" {
		return
	}
	s := &http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 30,
	}
	log.Println("listening", addr)
	log.Fatalln(s.ListenAndServe())
}
