package apiServerAction

import (
	"NKNMining/server/api/const"
	"NKNMining/server/api/response"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"NKNMining/storage"
	"NKNMining/crypto"
	"NKNMining/container"
)

const stopCmd = "stop"

var StopNodeAPI IRestfulAPIAction = &stopNodeAPI{}

type stopNodeAPI struct {
	restfulAPIBase
}

func (s *stopNodeAPI) URI(serverURI string) string {
	return serverURI + apiServerConsts.API_SERVER_URI_BASE + "/stop/node"
}

func (s *stopNodeAPI) Action(ctx *gin.Context) {
	response := apiServerResponse.New(ctx)

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

	if stopCmd != cmdStr {
		response.BadRequest(nil)
		return
	}

	container.Node.Stop()
	response.Success(nil)

	return
}