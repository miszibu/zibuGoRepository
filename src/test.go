package main

import (
	"zibuGoRepository/src/util/md5"
	"fmt"
)

func main(){
	// test md5 utils
	fmt.Println(md5.Md5PicByUrl("https://www.baidu.com/img/bd_logo1.png?where=super"))
}