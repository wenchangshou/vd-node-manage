package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
	"io/ioutil"
	"net/http"
)

func configSystemRoutes() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		req := make(map[string]interface{})
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &req)
		if err != nil {
			 AutoRender(w, nil, err)
			 return
		}
		server,ok:=req["server"]
		if !ok{
			AutoRender(w,nil,errors.New("server 必须填写"))
			return
		}
		code,ok:=req["code"]
		if !ok{
			AutoRender(w,nil,errors.New("code 必须填写"))
			return
		}
		rpcClient:=g.SingleConnRpcClient{RpcServer: fmt.Sprintf("%s:%d",server,8889)}


	})
}
