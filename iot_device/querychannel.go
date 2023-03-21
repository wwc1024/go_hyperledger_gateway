package iot_device

import (
	"fmt"
	_ "strings"

	_ "github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	_ "github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	_ "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	_ "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	_ "github.com/hyperledger/fabric-sdk-go/pkg/core/config/lookup"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	_ "github.com/pkg/errors"
)

var Flag = 0

//查询交易----未定义
/*func QueryTransInfo() {
	sdk, err := fabsdk.New(config.FromFile(configPath))
	if err != nil {
		panic(err)
	}
	clientContext := sdk.ChannelContext("mychannel", fabsdk.WithUser("Admin"))

	channelclient, err := ledger.New(clientContext)
	bci, err := channelclient.QueryInfo()
	if err != nil {
		fmt.Printf("failed to query for blockchain info: %v\n", err)
	}
	//当前区块高度和当前区块哈希
	if bci != nil {
		fmt.Printf("Retrieved ledger info: %v\n", bci)
	}

	//QueryBlock queries the ledger for Block by block number.
	block, err := channelclient.QueryBlock(9)
	if err != nil {
		fmt.Printf("failed to QueryBlock: %v\n", err)
	}
	if block != nil {
		fmt.Println("Retrieved block #9")
	}
	//QueryBlockByHash queries the ledger for block by block hash.
	block2, err := channelclient.QueryBlockByHash([]byte("\215\344\022\016/\315\3568\345\\\357\337\356\346\247m\257\227\235+\375\237\375\263\307W9k\263\016L;"))
	if err != nil {
		fmt.Printf("failed to query block by hash: %v\n", err)
	}

	if block2 != nil {
		fmt.Println("Retrieved block by hash")
	}

}*/

type ChannelOutput struct {
	Target     string
	Channel_id string
	Package_id string
}

//查询通道
func QueryChannel() ([]ChannelOutput, error) {
	//CreateRC
	sdk, err := fabsdk.New(config.FromFile(ConfigPath))
	if err != nil {
		fmt.Printf("Failed to create new cli SDK: %s", err)
		return nil, err
	}
	defer sdk.Close()
	_, err = sdk.Config()
	if err != nil {
		fmt.Printf("failed to sdk.Config():%s", err)
		return nil, err
	}
	clientContext := sdk.Context(fabsdk.WithUser(OrgAdmin), fabsdk.WithOrg(OrgName))
	resMgmtClient, err := resmgmt.New(clientContext)
	if err != nil {
		fmt.Printf("failed to query channel management client:%s", err)
		return nil, err
	}
	fmt.Println(Targets[0])
	//service
	outputList := []ChannelOutput{}
	channelQueryResponse, err := resMgmtClient.QueryChannels(
		resmgmt.WithTargetEndpoints(Targets[0]))
	if err != nil {
		fmt.Printf("QueryChannels return error: %s", err)
		return nil, err
	}

	if channelQueryResponse != nil {
		for _, channel := range channelQueryResponse.Channels {
			outputList = append(outputList, ChannelOutput{Target: Targets[0], Channel_id: channel.ChannelId})
			//若mychanneltwo通道存在，set Flag = 1
			if channel.ChannelId == ChannelId2 {
				Flag = 1
			}
			fmt.Printf("***  Channel :%v\n", channel.ChannelId)
		}
	}

	//reqPeers := resmgmt.WithTargetEndpoints(targets[0])
	//查看installedchaincode
	response, err := resMgmtClient.LifecycleQueryInstalledCC(resmgmt.WithTargetEndpoints(Targets[0]))
	if err != nil {
		fmt.Printf("failed to query installed chaincodes: %s\n", err)
		return nil, err
	}
	//查看commit情况
	// response, err := resMgmtClient.LifecycleQueryCommittedCC("mychannel", resmgmt.LifecycleQueryCommittedCCRequest{Name: "mycc"}, resmgmt.WithTargetEndpoints(targets[0]))

	if response != nil {
		fmt.Printf("Retrieved installed chaincodes:%v\n", response[0].PackageID)
		for i, _ := range response {
			outputList[i].Package_id = response[i].PackageID
		}
	}
	return outputList, nil
}

/*func orgTargetPeers(orgs []string, configBackend ...core.ConfigBackend) ([]string, error) {
	networkConfig := fab.NetworkConfig{}
	err := lookup.New(configBackend...).UnmarshalKey("organizations", &networkConfig.Organizations)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get organizations from config ")
	}

	var peers []string
	for _, org := range orgs {
		orgConfig, ok := networkConfig.Organizations[strings.ToLower(org)]
		if !ok {
			continue
		}
		peers = append(peers, orgConfig.Peers...)
		//fmt.Println(peers)
	}
	return peers, nil
}*/
