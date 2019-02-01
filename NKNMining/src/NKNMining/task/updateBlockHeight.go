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

	seedList := []string {
		"http://testnet-seed-0001.nkn.org:30003",
		"http://testnet-seed-0002.nkn.org:30003",
		"http://testnet-seed-0003.nkn.org:30003",
		"http://testnet-seed-0004.nkn.org:30003",
		"http://testnet-seed-0005.nkn.org:30003",
		"http://testnet-seed-0006.nkn.org:30003",
		"http://testnet-seed-0007.nkn.org:30003",
		"http://testnet-seed-0008.nkn.org:30003",
	}
	baseStart := uint64(1)
	mod := uint64(len(seedList))
	for {
		time.Sleep(10 * time.Second)
		baseStart += 1
		response, err := http.Post(seedList[baseStart % mod],
			"text/plain",
			strings.NewReader("{\"jsonrpc\":\"2.0\",\"id\":\"NKNMining\",\"method\":\"getlatestblockheight\",\"params\":{} }"))
		if nil != err {
			continue
		}

		resBody, err := ioutil.ReadAll(response.Body)
		if nil != err {
			response.Body.Close()
			continue
		}

		response.Body.Close()

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