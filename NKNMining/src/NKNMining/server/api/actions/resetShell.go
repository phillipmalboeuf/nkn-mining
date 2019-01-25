package apiServerAction

import (
	"NKNMining/server/api/const"
	"NKNMining/server/api/response"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"NKNMining/storage"
	"NKNMining/crypto"
)

const resetShellCmd = "resetShell"

var ResetShellAPI IRestfulAPIAction = &resetShellAPI{}

type resetShellAPI struct {
	restfulAPIBase
}

func (s *resetShellAPI) URI(serverURI string) string {
	return serverURI + apiServerConsts.API_SERVER_URI_BASE + "/reset/shell"
}

func (s *resetShellAPI) Action(ctx *gin.Context) {
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

	if resetShellCmd != cmdStr {
		response.BadRequest(nil)
		return
	}

	storage.NKNSetupInfo.Reset()

	response.Success(nil)

	return
}
