package nknReleaseQuery

import (
	"net/http"
	"NKNMining/common"
	"io/ioutil"
	"encoding/json"
	"runtime"
	"errors"
	"os"
	"io"
)

const (
	github_api_nkn_releases = "https://api.github.com/repos/nknorg/nkn/releases"
)

type GithubAssets struct {
	Name string `json:"name"`
	BrowserDownloadUrl string `json:"browser_download_url"`
}

type GithubReleaseInfo struct {
	TagName string `json:"tag_name"`
	Assets []GithubAssets `json:"assets"`
}

func LastVersion() (version  string, binUrl string, err error) {
	r, err := http.Get(github_api_nkn_releases)
	if nil != err {
		common.Log.Error(err)
		return
	}
	defer r.Body.Close()

	releaseInfo, err := ioutil.ReadAll(r.Body)
	if nil != err {
		common.Log.Error(err)
		return
	}

	var releaseArray = new([]GithubReleaseInfo)
	err = json.Unmarshal(releaseInfo, releaseArray)
	if nil != err {
		common.Log.Error(err," - ", string(releaseInfo))
		return
	}

	lastRelease := (*releaseArray)[0]
	version = lastRelease.TagName
	binZip := runtime.GOOS + "-" + runtime.GOARCH + ".zip"

	hasTargetBin := false
	for _, v := range lastRelease.Assets {
		if v.Name == binZip {
			binUrl = v.BrowserDownloadUrl
			hasTargetBin = true
			break
		}
	}

	if !hasTargetBin {
		err = errors.New("no such nkn bin file for your os-architecture")
		common.Log.Error(err)
		return
	}

	return
}


func DownloadNKN(url string, fileName string) (err error) {
	output, err := os.Create(fileName)
	if err != nil {
		common.Log.Error("Error while creating ", fileName, "-", err)
		return
	}
	defer output.Close()


	r, err := http.Get(url)
	if err != nil {
		common.Log.Error("Error while downloading ", url, "-", err)
		return
	}
	defer r.Body.Close()

	_, err = io.Copy(output, r.Body)
	if err != nil {
		common.Log.Error("Error while downloading ", url, "-", err)
		return
	}
	return
}

