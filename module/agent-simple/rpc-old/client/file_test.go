package rpcClient

import (
	"fmt"
	"testing"
)

func TestGetFileInfoByResourceReleaseID(t *testing.T) {
	server := InitFileServer("127.0.0.1:10051")
	fileInfo, err := server.GetFileInfoByProjectReleaseID("")
	fmt.Println(fileInfo, err)
}
