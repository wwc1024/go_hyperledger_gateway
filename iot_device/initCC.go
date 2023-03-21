package iot_device

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	_ "github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func ChannelInitCC() {
	sdk, err := fabsdk.New(config.FromFile("./config4.yaml"))
	if err != nil {
		panic(err)
	}

	// 获取访问 mychannel 通道的句柄，注意User 如果是合约的写入及查询，也可以使用User1，但创建通道等系统操作需要使用Admin
	clientContext := sdk.ChannelContext("mychanneltwo", fabsdk.WithUser("Admin"))

	// 创建channel客户端
	channelclient, err := channel.New(clientContext)
	if err != nil {
		panic(err)
	}
	req := channel.Request{
		ChaincodeID: "myccc",      //合约名字
		Fcn:         "InitLedger", //方法
		Args:        [][]byte{},
		IsInit:      true,
	}
	res, err := channelclient.Execute(req)
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed init")
	} else {
		fmt.Println("init success" + res.TransactionID)
	}
}
