package main

import (
	"os"
	"os/signal"
	"syscall"

	"go_gateway/golang_common/lib"
	"go_gateway/iot_device"
	_ "go_gateway/iot_device"
	"go_gateway/router"
)

//endpoint dashboard后台管理  server代理服务器
//config ./conf/prod/ 对应配置文件夹
// var config = flag.String("config", "", "input config file like ./conf/dev/")

func main() {
	// flag.Parse()
	// if *config == "" {
	// 	flag.Usage()
	// 	os.Exit(1)
	// }
	//开启获取dht数据循环

	go iot_device.DhtCronStart()

	//网络服务客户端
	lib.InitModule("./conf/dev/")
	defer lib.Destroy()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
