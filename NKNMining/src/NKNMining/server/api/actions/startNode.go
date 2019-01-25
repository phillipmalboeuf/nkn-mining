package apiServerAction

import (
	"NKNMining/server/api/const"
	"NKNMining/server/api/response"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"NKNMining/storage"
	"NKNMining/crypto"
	"NKNMining/container"
	"NKNMining/status"
)

const startCmd = "start"

var StartNodeAPI IRestfulAPIAction = &startNodeAPI{}

type startNodeAPI struct {
	restfulAPIBase
}

func (s *startNodeAPI) URI(serverURI string) string {
	return serverURI + apiServerConsts.API_SERVER_URI_BASE + "/start/node"
}

func (s *startNodeAPI) Action(ctx *gin.Context) {
	response := apiServerResponse.New(ctx)

	if !status.CanStartNode() {
		if status.IsChainDataDownloading() {
			response.Forbidden(nil)
		} else {
			response.Forbidden("invalid node status!")
		}
		return
	}

	inputJson, err := s.ExtractInput(ctx)
	if nil != err {
		response.BadRequest("invalid raw request!")
		return
	}

	basicData := &restfulAPIBaseRequest{}
	err = json.Unmarshal([]byte(inputJson), basicData)
	if nil != err {
		response.BadRequest("invalid request format!")
		return
	}

	cmdStr, err := crypto.AesDecrypt(basicData.Data,
		storage.NKNSetupInfo.GetRequestKey())
	if nil != err {
		response.BadRequest("invalid request data!")
		return
	}

	if startCmd != cmdStr {
		response.BadRequest(nil)
		return
	}

	wKey, err := crypto.AesDecrypt(storage.NKNSetupInfo.WKey, storage.NKNSetupInfo.GetWalletKey())
	if nil != err {
		response.InternalServerError(nil)
		return
	}

	_, err = container.Node.AsyncRun([]string{"-p", wKey, "--no-check-port"}, "")
	if nil != err {
		response.InternalServerError("start node failed!")
		return
	}

	response.Success(nil)

	return
}
