package util

import (
	"archive/zip"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
)

func UnZip(dst, src string) error {
	zr, err := zip.OpenReader(src)
	defer zr.Close()
	if err != nil {
		return err
	}
	if zr == nil {
		return errors.New("无法打开压缩文件")
	}
	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name)
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			continue
		}
		fr, err := file.Open()
		if err != nil {
			return err
		}
		fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		_, err = io.Copy(fw, fr)
		// 将解压的结果输出
		if err != nil {
			return err
		}
		fw.Close()
		fr.Close()
	}
	return nil
}
