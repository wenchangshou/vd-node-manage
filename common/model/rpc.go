package model

import "fmt"

type SimpleRpcResponse struct {
	Code int `json:"code"`
}
func (s *SimpleRpcResponse) String()string{
	return fmt.Sprintf("<Code: %d>",s.Code)
}

type NullRpcRequest struct{

}