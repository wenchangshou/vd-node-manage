package http

import (
	"encoding/json"
	"fmt"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/core/g"
	model2 "github.com/wenchangshou/vd-node-manage/module/core/g/model"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
		resp.Config.Server = server.(string)
		resp.Config.ID = resp.ID
		arr := strings.Split(server.(string), ":")
		if strings.Contains(resp.Config.Rpc.Address, "0.0.0.0") {
			resp.Config.Rpc.Address = strings.ReplaceAll(resp.Config.Rpc.Address, "0.0.0.0", arr[0])
		}
		if strings.Contains(resp.Config.Http.Address, "0.0.0.0") {
			resp.Config.Http.Address = strings.ReplaceAll(resp.Config.Http.Address, "0.0.0.0", arr[0])
		}
		g.StoreServerInfo(&resp.Config)
		RenderCustomMsgJson(w, 0, "success")
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		serverInfo := g.GetServerInfo()
		setupHeader(w, r)
		info := model2.ServerInfo{
			Name:     "主机",
			Register: serverInfo.Register,
			Address:  serverInfo.Server,
			Expired:  -1,
		}
		info.Detailed.Communication = true
		info.Detailed.Server = true
		RenderDataJson(w, info)
	})
	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		setupHeader(w, r)
		g.GetServerInfo()
		g.ResetServerInfo()
		RenderCustomMsgJson(w, 0, "success")
	})
}
