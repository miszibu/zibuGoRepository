package file

import (
	"os"
	"time"
	log "github.com/sirupsen/logrus"
	"runtime/pprof"
	"fmt"
)

//记录内存使用情况
func MemProfile() error {
	filePath := "F:/GoProject/zibuGoRepository/" + time.Now().Format("2006-01-02")
	file,err:=CreateFile(filePath)
	if err!=nil {
		return err
	}
	//runtime.GC() // 调用GC清理内存 可以查看GC后的内存使用情况
	if err := pprof.WriteHeapProfile(file); err != nil {
		log.Error("获取内存使用情况失败: ", err)
		return err
	}
	defer file.Close()
	return nil
}

/*
 * 创建文件 1.检查文件目录 2.创建文件
 * @param  filepath 文件目录
 * @return os.File 所创建文件指针 error
 */
func CreateFile(filePath string) (*os.File,error){
	//检查文件目录
	if err := CheckAndCreateFilePath(filePath); err != nil {
		log.Error("文件目录创建失败：", err)
		return &os.File{},err
	}
	//创建prof文件
	filePath = filePath+"/"+fmt.Sprintln("%s%s%%s",time.Now().Hour(),time.Now().Minute(),time.Now().Second())
	f, err := os.Create(filePath)
	if err != nil {
		log.Error("创建文件使用失败: ", err)
		return &os.File{},err
	}
	defer f.Close()
	return f,nil
}

/*
 * 检查文件夹是否存在 若不存在则创建该文件夹
 * @param filepath 文件目录
 * @return error：nil成功
 */
func CheckAndCreateFilePath(filePath string) (error) {
	filePath = filePath+time.Now().Format("2006-01-02")
	// 检查文件夹是否存在
	if _, err := os.Stat(filePath); err == nil {
		log.Infoln("文件目录存在:", filePath)
		return nil
	} else {
		log.Infoln("文件目录不存在:", filePath)
		err := os.MkdirAll(filePath, 0711)
		if err != nil {
			log.Infoln("文件目录创建失败",err)
			return err
		}
		log.Infoln("创建文件目录成功:",filePath)
		return nil
	}
}
