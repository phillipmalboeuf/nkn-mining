package apiServerAction

import (
	"NKNMining/server/api/const"
	"NKNMining/server/api/response"
	"github.com/gin-gonic/gin"
	"NKNMining/status"
	"NKNMining/task"
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
		response.Success(map[string] interface{} {
			"chainDataDownloadingProgress": downloadingProgress,
			"shellStatus": nsStatus,
			"blockHeight": task.CurrentHeight.Result,
			"nknNetworkHeight": task.TheNetworkHeight.Result,
			"syncStatus": task.NodeState.Result.SyncState,
			"nodeInfo": task.NodeState.Result,
			"neighbor": task.NodeNeighbor.Result,
		})
	}

	return
}
