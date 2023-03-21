package iot_device

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

//执行query的合约
func QueryAllCarsContent() ([]byte, error) {
	channelclient := FabsdkConfFile2()
	req := channel.Request{
		ChaincodeID: "myccc",        //合约名字
		Fcn:         "QueryAllCars", //方法
	}
	res, err := channelclient.Query(req)
	if err != nil {
		fmt.Println("initCC failed")
	} else {
		fmt.Println("initCC success")
	}
	fmt.Println(res.Payload)
	return res.Payload, err
}
