package storage

import (
	"io/ioutil"
	"NKNMining/common"
	"encoding/json"
)

const tempFile = "tempData.json"

type TempData struct {
	LastPPid int
}

var Temp = &TempData{}

func (t *TempData) Load() (err error) {
	tempData, err := ioutil.ReadFile(tempFile)

	if nil != err {
		return
	}

	err = json.Unmarshal(tempData, t)
	return
}

func (t *TempData) SaveParentPid(ppid int) (err error)  {
	t.LastPPid = ppid
	tempString, _ := json.MarshalIndent(t, "", "    ")

	err = ioutil.WriteFile(tempFile, tempString, 0666)
	if nil != err {
		common.Log.Fatal("save temp file failed!")
	}

	return
}
