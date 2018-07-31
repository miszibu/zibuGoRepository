package md5

import (
	"net/http"
	"io"
	log "github.com/sirupsen/logrus"
	"crypto/md5"
	"encoding/hex"
	"os"
	"math"
	"fmt"
)

 /**
  * 传入URL地址，获取文件流，对其进行MD5，返回MD5值
  * @param picUrl:url地址
  * @return string:文件MD5 error
  */
//todo：相关耗时计算代码 可注释
func Md5FileByUrl(picUrl string) (string,error)   {
	// get请求访问 URL图片
	res, err := http.Get(picUrl)
	if err != nil {
		log.Error("访问资源失败",err)
		return "",err
	}
	defer res.Body.Close()
	// 生成md5
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, res.Body); err != nil {
		log.Error("复制资源文件数据失败", err)
		return "",err
	}
	return hex.EncodeToString(md5hash.Sum([]byte(""))),nil
}

func aaaa(picUrl string) (string,error)   {
	// get请求访问 URL图片
	res, err := http.Get(picUrl)
	if err != nil {
		log.Error("访问资源失败",err)
		return "",err
	}
	defer res.Body.Close()
	// 生成md5
	md5hash := md5.New()

	if _, err := io.Copy(md5hash, res.Body); err != nil {
		log.Error("复制资源文件数据失败", err)
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
		log.Error("文件路径不存在:", localPath, err)
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

const filechunk = 8192 // we settle for 8KB
func bbbbbb(localPath string)(string , error){
	file, err := os.Open(localPath)

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	// calculate the file size
	info, _ := file.Stat()

	filesize := info.Size()

	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))

	hash := md5.New()

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))
		buf := make([]byte, blocksize)

		file.Read(buf)
		io.WriteString(hash, string(buf)) // append into the hash
	}

	fmt.Printf("%s checksum is %x\n", file.Name(), hash.Sum(nil))
	return hex.EncodeToString(hash.Sum([]byte(""))),nil
}