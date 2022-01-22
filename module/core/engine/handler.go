package engine

import (
	"encoding/json"
	"fmt"
	"github.com/wenchangshou/vd-node-manage/common/model"
)

// DeviceEvent 接收服务端事件
func (schedule *Schedule) DeviceEvent(_ string, message []byte) (r *model.EventReply, err error) {
	req := model.EventRequest{}
	reply := model.EventReply{}
	if err = json.Unmarshal(message, &req); err != nil {
		return
	}
	if req.Action == "openLayout" {
		reply = schedule.openLayout(req)
	} else if req.Action == "closeLayout" {
		reply = schedule.closeLayout(req)
	} else if req.Action == "control" {
		reply = schedule.control(req)
	} else if req.Action == "checkResourceExists" {
		reply = schedule.CheckResourceExists(req)
	} else if req.Action == "change" {
		reply = schedule.ChangeWindowSource(req)
		fmt.Println(string(req.Arguments), reply)
	}
	if !req.Reply {
		return nil, nil
	}
	return &reply, nil
}

// openLayout 打开布局
func (schedule *Schedule) openLayout(req model.EventRequest) model.EventReply {
	var (
		err error
	)
	reply := model.EventReply{EventID: req.EventID}
	args := model.OpenLayoutCmdParams{}
	if err := json.Unmarshal(req.Arguments, &args); err != nil {
		reply.Err = err
		reply.Msg = "解析json错误"
		return reply
	}
	if schedule.threadMap, err = schedule.layoutManage.OpenLayout(args); err != nil {
		reply.Err = err
		reply.Msg = "打开布局失败"
		return reply
	}
	return model.GenerateSimpleSuccessEventReply(req.EventID)
}
func (schedule *Schedule) closeLayout(req model.EventRequest) model.EventReply {
	var (
		err error
	)
	reply := model.EventReply{EventID: req.EventID}
	if err = schedule.layoutManage.CloseLayout(); err != nil {
		reply.Err = err
		reply.Msg = "关闭布局失败"
		return reply
	}
	return model.GenerateSimpleSuccessEventReply(req.EventID)
}
func (schedule *Schedule) control(req model.EventRequest) model.EventReply {
	reply := model.EventReply{
		EventID: req.EventID,
	}
	m := schedule.layoutManage
	args := model.ControlWindowCmdParams{}
	err := json.Unmarshal(req.Arguments, &args)
	if err != nil {
		reply.Err = err
		reply.Msg = "解析控制命令失败"
		return reply
	}
	err = m.Control(args, false)
	if err != nil {
		reply.Err = err
		reply.Msg = "控制窗口失败"
		return reply
	}
	return model.GenerateSimpleSuccessEventReply(req.EventID)
}
func (schedule Schedule) CheckResourceExists(req model.EventRequest) model.EventReply {
	return model.GenerateSimpleSuccessEventReply(req.EventID)
}

func (schedule Schedule) ChangeWindowSource(req model.EventRequest) model.EventReply {
	reply := model.EventReply{}
	args := model.OpenWindowCmdParams{}
	err := json.Unmarshal(req.Arguments, &args)
	if err != nil {
		reply.Err = err
		reply.Msg = "解析change失败"
		return reply
	}
	err = schedule.layoutManage.Change(args)
	if err != nil {
		reply.Err = err
		reply.Msg = "改变窗口源失败"
		return reply
	}
	return model.GenerateSimpleSuccessEventReply(req.EventID)
}
