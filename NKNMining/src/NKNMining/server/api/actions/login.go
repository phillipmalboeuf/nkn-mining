package apiServerAction

import (
	"NKNMining/server/api/const"
	"NKNMining/server/api/response"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"NKNMining/storage"
	"NKNMining/crypto"
)

var LoginAPI IRestfulAPIAction = &login{}

type loginData struct {
	Nonce 	string
	RandomKey string
}

type login struct {
	restfulAPIBase
}

func (l *login) URI(serverURI string) string {
	return serverURI + apiServerConsts.API_SERVER_URI_BASE + "/login"
}

func (l *login) Action(ctx *gin.Context) {
	response := apiServerResponse.New(ctx)
	if storage.SETUP_STEP_SUCCESS != storage.NKNSetupInfo.CurrentStep &&
	   storage.SETUP_STEP_GEN_WALLET != storage.NKNSetupInfo.CurrentStep {
		response.BadRequest("invalid request!")
		return
	}

	inputJson, err := l.ExtractInput(ctx)
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

	loginInfoJsonStr, err := crypto.AesDecrypt(basicData.Data, storage.NKNSetupInfo.Account + storage.NKNSetupInfo.Key)
	if nil != err {
		response.BadRequest("invalid request data!")
		return
	}

	loginInfo := &loginData{}
	err = json.Unmarshal([]byte(loginInfoJsonStr), loginInfo)
	if nil != err {
		response.BadRequest("invalid login data!")
		return
	}

	loginInfo.Nonce = basicData.PlainData

	if loginInfo.Nonce != basicData.PlainData {
		response.BadRequest("wrong login data!")
		return
	}


	wallet := &storage.Wallet{}
	wallet.Load()
	walletStrBytes, err := json.Marshal(wallet)

	responseData, _ := json.Marshal(map[string] string {
		"Nonce": loginInfo.Nonce,
		"SN": storage.NKNSetupInfo.SerialNumber,
		"Wallet": string(walletStrBytes),
	})

	responseStr, err := crypto.AesEncrypt(string(responseData), loginInfo.Nonce + loginInfo.RandomKey)

	if nil != err {
		response.InternalServerError(nil)
		return
	}

	if nil != err {
		response.InternalServerError(nil)
	} else {
		response.Success(responseStr)
	}
	return
}
