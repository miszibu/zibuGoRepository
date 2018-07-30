package md5

import (
	"net/http"
	"io"
	log "github.com/sirupsen/logrus"
	"crypto/md5"
	"encoding/hex"
	"os"
)

 /**
  * 传入URL地址，获取文件流，对其进行MD5，返回MD5值
  * @param picUrl:url地址
  * @return string:文件MD5 error
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
		log.Error("复制文件数据失败", err)
		return "",err
	}
	return hex.EncodeToString(md5hash.Sum([]byte(""))),nil
}

/*
 * 传入本地文件地址，获取文件流，对其进行MD5，返回MD5值
 * @param localPath:文件地址
 * @return string:文件MD5 error
 */
func Md5PicByLocalPath(localPath string)(string , error){
	// 查找目录
	if _, err := os.Stat(localPath); err != nil {
		log.Error("文件目录不存在:", localPath, err)
		return "",err
	} else {
		// 打开文件
		f,err:=os.Open(localPath)
		if err!=nil {
			log.Error("打开文件失败:", localPath, err)
			return "",err
		}
		//生成md5
		md5hash := md5.New()
		if _, err := io.Copy(md5hash, f); err != nil {
			log.Error("复制文件数据失败", err)
			return "",err
		}
		return hex.EncodeToString(md5hash.Sum([]byte(""))),nil
	}
}


