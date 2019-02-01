package apiServerAction

import (
	"NKNMining/server/api/const"
	"NKNMining/server/api/response"
	"github.com/gin-gonic/gin"
	"NKNMining/status"
	"NKNMining/task"
	"NKNMining/storage"
)

var GetStatusAPI IRestfulAPIAction = &getStatus{}

type getStatus struct {
	restfulAPIBase
}

func (g *getStatus) URI(serverURI string) string {
	return serverURI + apiServerConsts.API_SERVER_URI_BASE + "/status"
}

func (g *getStatus) Action(ctx *gin.Context) {
	response := apiServerResponse.New(ctx)
	nsStatus, errInfo := status.GetServerStatus()

	if "" != errInfo {
		response.InternalServerError(errInfo)
	} else {
		downloadingProgress := 100.0
		if status.IsChainDataDownloading() {
			downloadingProgress = task.GetChainDataDownloadProgress()
		}
		fromLocal := ctx.Params.ByName(apiServerConsts.PARAM_FROM_LOCAL)
		result := map[string] interface{} {
			"chainDataDownloadingProgress": downloadingProgress,
			"shellStatus": nsStatus,
			"blockHeight": task.CurrentHeight.Result,
			"nknNetworkHeight": task.TheNetworkHeight.Result,
			"syncStatus": task.NodeState.Result.SyncState,
			"nodeInfo": task.NodeState.Result,
			"neighbor": task.NodeNeighbor.Result,
		}

		if "" != fromLocal {
			result["sn"] = storage.NKNSetupInfo.SerialNumber
		}

		response.Success(result)
	}

	return
}
