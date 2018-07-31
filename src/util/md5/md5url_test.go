package md5

import (
	"testing"
	"path/filepath"
	"time"
	"github.com/sirupsen/logrus"
)

func TestMd5PicByLocalPath(t *testing.T) {
	beginTime := time.Now()
	//localpath := filepath.Join("F:\\typora-setup-x64.exe")
	//localpath := filepath.Join("F:\\《区块链技术指南》超清彩图版PDF.pdf")
	//localpath := filepath.Join("F:\\android-sdk_r24.4.1-windows.zip")
	localpath := filepath.Join("F:\\md5.java")
	logrus.Infoln(Md5PicByLocalPath(localpath))

	logrus.Infoln(bbbbbb(localpath))
	logrus.Info(time.Since(beginTime))
}

func TestMd5FileByUrl(t *testing.T) {
	beginTime := time.Now()
	//这个是不一致的
	//localpath :="https://dl.google.com/tag/s/appguid%3D%7B8A69D345-D564-463C-AFF1-A69D9E530F96%7D%26iid%3D%7B09019EAA-9A86-C093-EAB6-481B14A38FD2%7D%26lang%3Dzh-CN%26browser%3D4%26usagestats%3D0%26appname%3DGoogle%2520Chrome%26needsadmin%3Dprefers%26ap%3Dx64-stable-statsdef_1%26installdataindex%3Ddefaultbrowser/update2/installers/ChromeSetup.exe"
	localpath := "https://d11.baidupcs.com/file/a89f4a5e29be8edc59a0df3b05eee2c6?bkt=p3-1400a89f4a5e29be8edc59a0df3b05eee2c64c8166fd000000010105&xcode=26702497e3539c19bee1b69fe8fc0c2e84eff2bb8d8f2be58a4824fc1ec097a8546e846baceda86bb0b3b3d56a3e78879a7e3ac4ae9d7ad8&fid=3344747512-250528-574136899459084&time=1532999597&sign=FDTAXGERLQBHSK-DCb740ccc5511e5e8fedcff06b081203-KDJ6uNy5FLc%2FtowcUp2qt0vFefM%3D&to=d11&size=65797&sta_dx=65797&sta_cs=410&sta_ft=zip&sta_ct=7&sta_mt=7&fm2=MH%2CNanjing02%2CAnywhere%2C%2Czhejiang%2Cct&resv0=cdnback&resv1=0&vuk=2556959591&iv=0&newver=1&newfm=1&secfm=1&flow_ver=3&pkey=1400a89f4a5e29be8edc59a0df3b05eee2c64c8166fd000000010105&sl=76480590&expires=8h&rt=sh&r=606468315&mlogid=4902596884539110553&vbdid=1111190313&fin=%E3%80%90%E5%87%AF%E6%BA%90%E3%80%91%E5%AF%86%E7%A0%81+BY%EF%BC%9A%E5%92%95%E5%98%9F%E5%92%95%E5%99%9C%E5%90%90%E6%B3%A1%E6%B3%A1.zip&fn=%E3%80%90%E5%87%AF%E6%BA%90%E3%80%91%E5%AF%86%E7%A0%81+BY%EF%BC%9A%E5%92%95%E5%98%9F%E5%92%95%E5%99%9C%E5%90%90%E6%B3%A1%E6%B3%A1.zip&rtype=1&dp-logid=4902596884539110553&dp-callid=0.1.1&hps=1&tsl=80&csl=80&csign=jzX1HVXBpTEPjS6DuoTvLo%2Bj7OI%3D&so=0&ut=6&uter=4&serv=0&uc=1787862661&ic=690492119&ti=26fa64dbec28822470ab60547d1263e072f39da5355341c8305a5e1275657320&by=themis"
	//localpath := "https://www.baidu.com/img/bd_logo1.png?where=super"
	md5,err:= Md5FileByUrl(localpath)
    if err!=nil{
    	t.Error("md5文件失败",err)
	}
	logrus.Infoln("文件MD5值",md5)
	logrus.Info("MD5文件总共用时",time.Since(beginTime))
}