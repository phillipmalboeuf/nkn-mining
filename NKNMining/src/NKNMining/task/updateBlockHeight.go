package task

import (
	"NKNMining/network/chainDataTypes/rpcApiResponse"
	"NKNMining/status"
	"NKNMining/common"
	"time"
	"NKNMining/network/rpcRequest"
	"NKNMining/network"
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
)

var CurrentHeight = &rpcApiResponse.BlockHeight{}
var TheNetworkHeight = &rpcApiResponse.BlockHeight{}

func UpdateNetworkHeight() {
	for {
		time.Sleep(5 * time.Second)

		response, err := http.Post("http://testnet-node-0001.nkn.org:30003",
			"text/plain",
			strings.NewReader("{\"jsonrpc\":\"2.0\",\"id\":\"NKNMining\",\"method\":\"getlatestblockheight\",\"params\":{} }"))
		if nil != err {
			continue
		}

		resBody, err := ioutil.ReadAll(response.Body)
		if nil != err {
			continue
		}

		json.Unmarshal(resBody, TheNetworkHeight)
	}
}

func UpdateCurrentHeight() {
	for {
		shellStatus, errInfo := status.GetServerStatus()

		if "" != errInfo || shellStatus != common.NS_STATUS_NODE_RUNNING {
			if "" != errInfo {
				common.Log.Error(errInfo)
			}
			time.Sleep(5 * time.Second)
			continue
		}

		ret, err := rpcRequest.Api.Call(network.RPC_API_BLOCK_HEIGHT, nil, false, 3)
		if nil == err {
			CurrentHeight = ret.(*rpcApiResponse.BlockHeight)
		}

		time.Sleep(5 * time.Second)
	}
}