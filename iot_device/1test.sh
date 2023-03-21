#!/bin/bash
echo "now:{name: $1 ;temp: $2 ;hum: $3 }"
BACKUP=/home/neo/go/src/github.com/hyperledger/fabric-samples/multiple-deployment/querydht/query
#DATETIME=$(date +%Y-$m-%d_%H%M%S)
name="$1"
temp="$2"
hum="$3"

docker exec -i cli1 peer chaincode invoke -o orderer0.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer0.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses peer0.org2.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"Args":["CreateDht","'$name'","'$temp'","'$hum'"]}' --waitForEvent >& ${BACKUP}/backup.txt
echo "Dates :"
cat  ${BACKUP}/backup.txt

