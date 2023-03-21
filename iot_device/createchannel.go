package iot_device

import (
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/test/metadata"
)

/*func main() {
	rc := FabsdkConfFile()
	// Peer joins channel 'mychannel'createChannel()
	//createChannel(rc)
	join(rc)

}*/

func CreateChannel() {
	sdk, err := fabsdk.New(config.FromFile(ConfigPath))
	if err != nil {
		panic(err)
	}
	clientContext := sdk.Context(fabsdk.WithUser(OrgAdmin), fabsdk.WithOrg(OrgName))

	rc, err := resmgmt.New(clientContext)
	if err != nil {
		fmt.Printf("failed to New rc %v\n", err)
	}
	// Read channel configuration tx
	channelConfigTxPath := filepath.Join(metadata.GetProjectPath(), metadata.ChannelConfigPath, ChannelConfigTxFile)
	r, err := os.Open(channelConfigTxPath)
	if err != nil {
		fmt.Printf("failed to open channel config: %s\n", err)
	}
	defer r.Close()

	// Create new channel 'mychanneltwo'
	resp, err := rc.SaveChannel(resmgmt.SaveChannelRequest{ChannelID: ChannelId2, ChannelConfig: r}, resmgmt.WithOrdererEndpoint(OrderName))
	if err != nil {
		fmt.Printf("failed to save channel: %s\n", err)
	}
	if resp.TransactionID == "" {
		fmt.Println("Failed to save channel")
	} else {
		fmt.Println("Saved channel")
	}
}

func JoinChannel(rc *resmgmt.Client) {
	err := rc.JoinChannel(ChannelId2, resmgmt.WithTargetEndpoints(PeerName), resmgmt.WithOrdererEndpoint(OrderName))
	if err != nil {
		fmt.Printf("failed to join channel: %s\n", err)
	}
	fmt.Println(">> 加入通道成功")
}
