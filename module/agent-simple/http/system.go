package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
		server, ok := req["server"]
		if !ok {
			AutoRender(w, nil, errors.New("server 必须填写"))
			return
		}
		code, ok := req["code"]
		if !ok {
			AutoRender(w, nil, errors.New("code 必须填写"))
			return
		}
		rpcClient := &g.SingleConnRpcClient{
			RpcServer: fmt.Sprintf("%s", server),
			Timeout:   time.Second,
		}
		r3 := model.SimpleRpcResponse{}
		err = rpcClient.Call("Device.Ping", &model.NullRpcRequest{}, &r3)
		if err != nil {
			AutoRender(w, nil, errors.New("服务端通讯异常"))
			return
		}
		var resp *model.DeviceRegisterResponse
		r2 := &model.DeviceRegisterRequest{
			HardwareCode: g.Hardware().ID,
			Code:         code.(string),
			ConnType:     "link",
		}
		err = rpcClient.Call("Device.Register", r2, &resp)
		if err != nil || resp.Code != 0 {
			log.Println("call Device.Request fail:", err, "Request:", req, "Response:", resp)
			AutoRender(w, nil, errors.New("注册失败:"+resp.Msg))
			return
		}
		g.SetRegisterStatus(true, resp.ID, resp.HttpAddress, resp.RpcAddress, resp.RedisAddress)
		RenderMsgJson(w, "success")
	})
	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		g.SetRegisterStatus(false, 0, "", "", "")
	})
}
