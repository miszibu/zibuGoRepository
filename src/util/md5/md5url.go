package md5

import (
	"net/http"
	"io"
	log "github.com/sirupsen/logrus"
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

/* 
* 传入URL地址，获取文件流，对其进行MD5，返回MD5值
* @param url地址
* @return 文件MD5 error
*/

func Md5PicByUrl(picUrl string) (string,error)   {
	// get请求访问 URL图片
	res, err := http.Get(picUrl)
	if err != nil {
		log.Error("访问图片失败",err)
		return "",err
	}
	defer res.Body.Close()
	// 生成md5
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, res.Body); err != nil {
		fmt.Println("Copy", err)
		return "",err
	}
	return hex.EncodeToString(md5hash.Sum([]byte(""))),nil
}


