package md5

import (
	"testing"
	"path/filepath"
	"time"
	"github.com/sirupsen/logrus"
)

func TestMd5PicByLocalPath(t *testing.T) {
	beginTime := time.Now()
	localpath := filepath.Join("F:\\typora-setup-x64.exe")
	//localpath := filepath.Join("F:\\《区块链技术指南》超清彩图版PDF.pdf")
	//localpath := filepath.Join("F:\\android-sdk_r24.4.1-windows.zip")
	//localpath := filepath.Join("F:\\md5.java")
	logrus.Infoln(Md5PicByLocalPath(localpath))
	logrus.Info(time.Since(beginTime))
}

func TestMd5PicByUrl(t *testing.T) {
	beginTime := time.Now()
	localpath := "https://www.baidu.com/img/bd_logo1.png?where=super"
	Md5PicByUrl(localpath)
	logrus.Infoln(Md5PicByUrl(localpath))
	logrus.Info(time.Since(beginTime))
}