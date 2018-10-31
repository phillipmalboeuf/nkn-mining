package task

import (
	"NKNMining/network/rpcRequest"
	"NKNMining/network"
	"NKNMining/network/chainDataTypes/rpcApiResponse"
	"time"
	"NKNMining/status"
	"NKNMining/common"
)

var NodeState = &rpcApiResponse.NodeInfo{}

func UpdateNodeState() {
	for {
		shellStatus, errInfo := status.GetServerStatus()

		if "" != errInfo || shellStatus != common.NS_STATUS_NODE_RUNNING {
			time.Sleep(5 * time.Second)
			continue
		}

		ret, err := rpcRequest.Api.Call(network.RPC_API_NODE_STATE, nil, false, 3)
		if nil == err {
			NodeState = ret.(*rpcApiResponse.NodeInfo)
		}

		time.Sleep(5 * time.Second)
	}
}