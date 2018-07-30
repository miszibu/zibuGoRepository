package main

import (
	"fmt"
	"time"
)

func main(){

	// test md5 utils
	//fmt.Println(md5.Md5PicByUrl("https://www.baidu.com/img/bd_logo1.png?where=super"))

	// test file CheckAndCreateFilePath
	/*if err:=file.MemProfile();err!=nil{
		fmt.Println("error",err)
	}else{
		fmt.Println("success",err)
	}*/
	a,b,c :=time.Now().Clock()
	str1 := fmt.Sprintf("%d:%d:%d",time.Now().Hour(),time.Now().Second(),time.Now().Second())
	fmt.Println(a,b,c,str1)
}