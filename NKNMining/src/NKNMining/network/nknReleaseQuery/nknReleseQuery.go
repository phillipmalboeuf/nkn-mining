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
	"time"
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

func checkDownloaded(downloadedBytes *float64, complete *bool, output *os.File) {
	if nil == downloadedBytes {
		return
	}

	for  {
		time.Sleep(time.Millisecond * 50)
		fi, err := output.Stat()
		if nil != err {
			common.Log.Error(err)
			break
		}

		*downloadedBytes = float64(fi.Size())
		if *complete {
			break
		}
	}
}


func DownloadNKN(url string, fileName string, downloadedBytes *float64) (err error) {
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

	complete := false
	if nil != downloadedBytes {
		go checkDownloaded(downloadedBytes, &complete, output)
	}

	_, err = io.Copy(output, r.Body)
	complete = true
	if err != nil {
		common.Log.Error("Error while downloading ", url, "-", err)
		return
	}
	return
}

