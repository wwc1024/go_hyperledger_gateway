package iot_device

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

const (
	//ChannelId           = "mychanneltwo"
	ChannelId1 = "mychannel"
	ConfigPath = "./config2.yaml"

	ConfigPath2         = "./config4.yaml"
	ChannelId2          = "mychanneltwo"
	CcID                = "myccc"
	CcVersion           = "0"
	Ccpath              = "fabcar"
	ChannelConfigTxFile = "../../../../../../../channel-artifacts/channeltwo.tx"
	OrgName             = "Org1"
	OrgMsp              = "Org1MSP"
	OrgAdmin            = "Admin"
	OrderName           = "orderer0.example.com"
	PeerName            = "peer0.org1.example.com"
	Countdhtnumber      = 0

	EndorsementPlugin = "escc"
	ValidationPlugin  = "vscc"
)

var (
	Targets         = []string{"peer0.org1.example.com"}
	Approvesequence = int64(1)
	Querysequence   = int64(1)
	Commitsequence  = int64(2)
)

//定义解析后结构体
type Dht struct {
	Id          string `json:"id"`
	DhtId       string `json:"dhtId"`
	Now         string `json:"time"`
	Temperature string `json:"temperature"`
	Humitidy    string `json:"humitidy"`
}

//访问 mychannel 通道
func FabsdkConfFile() *channel.Client {
	sdk, err := fabsdk.New(config.FromFile("./iot_device/config2.yaml"))
	if err != nil {
		panic(err)
	}

	// 获取访问 mychannel 通道的句柄，注意User 如果是合约的写入及查询，也可以使用User1，但创建通道等系统操作需要使用Admin
	clientContext := sdk.ChannelContext("mychannel", fabsdk.WithUser("Admin"))

	// 创建channel客户端
	channelclient, err := channel.New(clientContext)
	if err != nil {
		panic(err)
	}
	return channelclient
}

//访问 mychanneltwo 通道
func FabsdkConfFile2() *channel.Client {
	sdk, err := fabsdk.New(config.FromFile("./iot_device/config4.yaml"))
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
	return channelclient
}

func FabsdkConfFileRC() *resmgmt.Client {
	sdk, err := fabsdk.New(config.FromFile(ConfigPath))
	if err != nil {
		panic(err)
	}
	clientContext := sdk.Context(fabsdk.WithUser(OrgAdmin), fabsdk.WithOrg(OrgName))

	rc, err := resmgmt.New(clientContext)
	if err != nil {
		fmt.Printf("failed to New rc %v\n", err)
	}
	return rc
}
