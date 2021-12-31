package http

import (
	"encoding/json"
	"fmt"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/agent-simple/g"
	model2 "github.com/wenchangshou/vd-node-manage/module/agent-simple/g/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func configSystemRoutes() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		setupHeader(w, r)
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
			RenderCustomMsgJson(w, 400, "server 必须填写")
			return
		}
		code, ok := req["code"]
		if !ok {
			RenderCustomMsgJson(w, 400, "code 必须填写")
			return
		}
		rpcClient := &g.SingleConnRpcClient{
			RpcServer: fmt.Sprintf("%s", server),
			Timeout:   time.Second,
		}
		r3 := model.SimpleRpcResponse{}
		err = rpcClient.Call("Device.Ping", &model.NullRpcRequest{}, &r3)
		if err != nil {
			RenderCustomMsgJson(w, 400, "服务端通讯异常:"+err.Error())
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
			RenderCustomMsgJson(w, 400, "注册失败:"+resp.Msg)
			return
		}
		g.SetRegisterStatus(true, server.(string), resp.ID, resp.HttpAddress, resp.RpcAddress, resp.RedisAddress)
		RenderCustomMsgJson(w, 0, "success")
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		setupHeader(w, r)
		cfg := g.Config()
		info := model2.ServerInfo{
			Name:     "主机",
			Register: cfg.Server.Register,
			Address:  cfg.Server.Address,
			Expired:  -1,
		}
		info.Detailed.Communication = true
		info.Detailed.Server = true
		//rtu["register"] = cfg.Server.Register
		//rtu["address"] = cfg.Server.Address
		//rtu["expired"] = -1
		//rtu["detailed"] = struc

		RenderDataJson(w, info)
	})
	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		setupHeader(w, r)
		g.SetRegisterStatus(false, "", 0, "", "", "")
		RenderCustomMsgJson(w, 0, "success")
	})
}
