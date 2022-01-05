package model

type DownloadStatus int

const (
	Downloading DownloadStatus = iota
	DownloadErr
	DownloadDone
)

type TaskDownloadInfo struct {
	BytesComplete int64   `json:"bytesComplete"`
	Size          int64   `json:"size"`
	Process       float64 `json:"process"`
}
