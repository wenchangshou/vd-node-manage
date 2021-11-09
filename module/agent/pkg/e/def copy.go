package e
//type RequestCmd struct{
//	Action string
//
//}
type Event struct{
	Gid string `json:"gid"`
}
var (
	DOWNLOAD_WAITING=0//等待下载
	DOWNLOAD_ERROR=-1//下载错误
	DOWNLOADING=1 //下载中
	DOWNLOAD_COMPLETE=2//下载完成
	FILE_VERIFICATION_FIELD=10005//文件检验失败
	TYPE_APPLICATION=20001
	TYPE_RESOURCE=20002
	UPDATE_MODEL_INCREMENT=30001
	UPDATE_MODEL_FULL=30002
)
type RequestCmd struct{
	MessageType string `json:"messageType"`
	SocketName string `json:"socketName"`
	SocketType string `json:"SocketType"`
	Arguments map[string]interface{} `json:"Arguments"`

}