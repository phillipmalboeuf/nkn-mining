package storage

import (
	"io/ioutil"
	"NKNMining/common"
	"bytes"
	"encoding/json"
	"github.com/pborman/uuid"
	"fmt"
)

const (
	setupFile = "setup/info.json"
)

const (
	SETUP_STEP_INIT = iota
	SETUP_STEP_CREATE_ACCOUNT
	SETUP_STEP_GEN_WALLET
	SETUP_STEP_SUCCESS
	SETUP_NODE_UPDATE
)

func InitSetupInfo()  {
	NKNSetupInfo.Load()

	if SETUP_STEP_CREATE_ACCOUNT == NKNSetupInfo.CurrentStep {
		fmt.Println("serial number(sn): " + NKNSetupInfo.SerialNumber)
		return
	}

	if SETUP_STEP_INIT != NKNSetupInfo.CurrentStep {
		return
	}

	if "" == NKNSetupInfo.SerialNumber {
		sn := "NKN-" + uuid.NewUUID().String()
		NKNSetupInfo.SerialNumber = sn
	}

	NKNSetupInfo.CurrentStep = SETUP_STEP_CREATE_ACCOUNT
	if nil == NKNSetupInfo.Save() {
		fmt.Println("serial number(sn): " + NKNSetupInfo.SerialNumber)
	} else {
		common.Log.Fatal("initialization NKN setup info failed!")
	}

	NKNSetupInfo.Save()
}

type SetupInfo struct {
	SerialNumber string
	CurrentStep	int
	Account string
	Key string
	WKey string
	BinVersion string
	SelfNode string
}

var NKNSetupInfo = &SetupInfo{}

func (s *SetupInfo) Reset()  {
	NKNSetupInfo = &SetupInfo{SerialNumber: NKNSetupInfo.SerialNumber, CurrentStep: SETUP_STEP_CREATE_ACCOUNT}
	NKNSetupInfo.Save()
}

func (s *SetupInfo) GetRequestKey() string {
	return s.SerialNumber + s.Key
}


func (s *SetupInfo) GetWalletKey() string {
	return s.SerialNumber + s.Key + s.Account
}

func (s *SetupInfo) Load()  {
	setupInfo, err := ioutil.ReadFile(setupFile)
	if nil != err {
		common.Log.Fatal("no setup info!")
	}

	setupInfo = bytes.TrimPrefix(setupInfo, []byte("\xef\xbb\xbf"))

	err = json.Unmarshal(setupInfo, s)
	if nil != err {
		common.Log.Fatal("setup file damaged!")
	}
}

func (s *SetupInfo) Save() error {
	setupInfo, _ := json.MarshalIndent(s, "", "    ")

	err := ioutil.WriteFile(setupFile, setupInfo, 0666)
	if nil != err {
		common.Log.Fatal("save setup file failed!")
	}

	return err
}