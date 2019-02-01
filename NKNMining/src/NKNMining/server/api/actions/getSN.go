package apiServerAction

import (
	"NKNMining/server/api/const"
	"NKNMining/server/api/response"
	"github.com/gin-gonic/gin"
	"NKNMining/storage"
)

var GetSNAPI IRestfulAPIAction = &getSNAPI{}

type getSNAPI struct {
	restfulAPIBase
}

func (g *getSNAPI) URI(serverURI string) string {
	return serverURI + apiServerConsts.API_SERVER_URI_BASE + "/get/sn"
}

func (g *getSNAPI) Action(ctx *gin.Context) {
	response := apiServerResponse.New(ctx)

	isLocalRequest := ctx.Params.ByName(apiServerConsts.PARAM_FROM_LOCAL)
	if "" == isLocalRequest {
		response.Forbidden(nil)
		return
	}

	response.Success(storage.NKNSetupInfo.SerialNumber)
	return
}
