package iot_device

import (
	"fmt"
	"path/filepath"

	pb "github.com/hyperledger/fabric-protos-go/peer"
	_ "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	_ "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	lcpackager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/lifecycle"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/test/metadata"
	"github.com/hyperledger/fabric/common/policydsl"
)

/*
const (
	ChannelId           = "mychanneltwo"
	ChannelId1           = "mychannel"
	ChannelId2           = "mychanneltwo"
	ConfigPath          = "./config2.yaml"
	CcID                = "myccc"
	CcVersion           = "0"
	Ccpath              = "fabdht2"
	ChannelConfigTxFile = "../../../../../../../channel-artifacts/channeltwo.tx"
	OrgName             = "Org1"
	OrgMsp              = "Org1MSP"
	OrgAdmin            = "Admin"
	OrderName           = "orderer0.example.com"
	PeerName            = "peer0.org1.example.com"

	Approvesequence     = 1
	EndorsementPlugin   = "escc"
	ValidationPlugin    = "vscc"
	Querysequence       = 1
	Commitsequence      = 1
)
var Targets =  []string{"peer0.org1.example.com"}
*/

func InstallCC() error {
	// sdk, err := fabsdk.New(config.FromFile(ConfigPath))
	//在mychanneltwo安装链码时
	sdk, err := fabsdk.New(config.FromFile(ConfigPath2))
	if err != nil {
		fmt.Printf("installCC newsdk failed,%v", err)
		return err
	}
	rcp := sdk.Context(fabsdk.WithUser(OrgAdmin), fabsdk.WithOrg(OrgName))
	rc, err := resmgmt.New(rcp)
	if err != nil {
		fmt.Printf("failed to create resource client: %v", err)
		return err
	}

	ccPolicy := policydsl.SignedByAnyMember([]string{OrgMsp})
	// pack the chaincode
	/*ccPkg, err := gopackager.NewCCPackage(ccpath, "/home/json/sdk/go")
	if err != nil {
		fmt.Printf( "pack chaincode error, %v",err)
	}*/
	label := CcID + "_" + CcVersion // 链码的标签
	ccPath := filepath.Join(metadata.GetProjectPath(), Ccpath)
	desc := &lcpackager.Descriptor{
		Path:  ccPath,
		Type:  pb.ChaincodeSpec_GOLANG,
		Label: label,
	}
	ccPkg, err := lcpackager.NewCCPackage(desc)
	if err != nil {
		fmt.Printf("lcpackager failed,%v\n", err)
		return err
	}
	//fmt.Printf("%v,%v\n", desc.Label, ccPkg)

	// new request of installing chaincode
	installCCReq := resmgmt.LifecycleInstallCCRequest{
		Label:   label,
		Package: ccPkg,
	}

	reqPeers := resmgmt.WithTargetEndpoints(PeerName)
	packageID := lcpackager.ComputePackageID(installCCReq.Label, installCCReq.Package)
	responses, err := rc.LifecycleInstallCC(installCCReq, reqPeers)
	if err != nil {
		fmt.Printf("failed to install chaincode: %v\n", err)
		return err
	}
	if len(responses) > 0 {
		fmt.Printf("Chaincode installed,%v\n", packageID)
	}

	//------不走instantiateCC，走打包、安装、审批、提交--------
	//approve
	/*org1Peers, err := integration.DiscoverLocalPeers(mc.org1AdminClientContext, 2)
	require.NoError(t, err)
	org2Peers, err := integration.DiscoverLocalPeers(mc.org2AdminClientContext, 2)
	require.NoError(t, err)
	ccPolicy := policydsl.SignedByNOutOfGivenRole(2, mb.MSPRole_MEMBER, []string{"Org1MSP", "Org2MSP"})*/

	approveCCReq := resmgmt.LifecycleApproveCCRequest{
		Name:              CcID,
		Version:           CcVersion,
		PackageID:         packageID,
		Sequence:          Approvesequence,
		EndorsementPlugin: EndorsementPlugin,
		ValidationPlugin:  ValidationPlugin,
		SignaturePolicy:   ccPolicy,
		InitRequired:      true,
	}

	txnID, err := rc.LifecycleApproveCC(ChannelId2, approveCCReq, resmgmt.WithTargetEndpoints(PeerName), resmgmt.WithOrdererEndpoint(OrderName))
	if err != nil {
		fmt.Printf("LifecycleApproveCC failed:%v\n", err)
		return err
	} else {
		fmt.Println(txnID)
		Approvesequence++
	}

	//queryApproveCC有问题

	// queryApprovedCCReq := resmgmt.LifecycleQueryApprovedCCRequest{
	// 	Name:     ChannelId2,
	// 	Sequence: Querysequence,
	// }
	// resp, err := rc.LifecycleQueryApprovedCC(ChannelId2, queryApprovedCCReq, resmgmt.WithTargetEndpoints(PeerName))
	// if err != nil {
	// 	fmt.Printf("LifecycleQueryApprovedCC failed:%v\n", err)
	// 	return err
	// } else {
	// 	fmt.Println(resp)
	// 	Querysequence++
	// }

	//commit
	commitReq := resmgmt.LifecycleCommitCCRequest{
		Name:              CcID,
		Version:           CcVersion,
		Sequence:          Commitsequence,
		EndorsementPlugin: EndorsementPlugin,
		ValidationPlugin:  ValidationPlugin,
		SignaturePolicy:   ccPolicy,
		InitRequired:      true,
	}
	txnID2, err := rc.LifecycleCommitCC(ChannelId2, commitReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithTargetEndpoints(PeerName), resmgmt.WithOrdererEndpoint(OrderName))
	if err != nil {
		fmt.Printf("LifecycleCommitCC failed: %v\n", err)
		return err
	} else {
		fmt.Println(txnID2)
		Commitsequence++
	}

	// queryCommit有问题
	// qcommitReq := resmgmt.LifecycleQueryCommittedCCRequest{
	// 	Name: CcID,
	// }
	// resp2, err := rc.LifecycleQueryCommittedCC(ChannelId2, qcommitReq, resmgmt.WithTargetEndpoints(PeerName), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// 	return err
	// } else {
	// 	fmt.Println(resp2)
	// }

	return nil
}
