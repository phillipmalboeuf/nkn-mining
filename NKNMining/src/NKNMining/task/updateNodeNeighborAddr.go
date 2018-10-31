package task

import (
	"NKNMining/network/chainDataTypes/rpcApiResponse"
	"NKNMining/status"
	"NKNMining/common"
	"time"
	"NKNMining/network/rpcRequest"
	"NKNMining/network"
)

var NodeNeighbor = &rpcApiResponse.NodeAddr{}

func UpdateNodeNeighbor() {
	for {
		shellStatus, errInfo := status.GetServerStatus()

		if "" != errInfo || shellStatus != common.NS_STATUS_NODE_RUNNING {
			time.Sleep(5 * time.Second)
			continue
		}

		ret, err := rpcRequest.Api.Call(network.RPC_API_NODE_NEIGHBOR_ADDR, nil, false, 3)
		if nil == err {
			NodeNeighbor = ret.(*rpcApiResponse.NodeAddr)
		}

		time.Sleep(5 * time.Second)
	}
}