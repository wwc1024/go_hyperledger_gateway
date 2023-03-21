package iot_device

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

//执行query的合约
func QueryAllDhtsContent() ([]byte, error) {
	channelclient := FabsdkConfFile()
	req := channel.Request{
		ChaincodeID: "mycc",         //合约名字
		Fcn:         "QueryAllDhts", //方法
	}
	res, err := channelclient.Query(req)
	return res.Payload, err
}
