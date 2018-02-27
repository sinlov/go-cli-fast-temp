package aubnig

import (
	"path/filepath"
	"fmt"
	"os"
	"strings"
	sFiles "github.com/sinlov/golang_utils/files"
	sCli "github.com/sinlov/golang_utils/cli"
	"errors"
)

const (
	defaultConfPath = "temp.conf"
	gitRepo         = "go-cli-fast-temp"
	gitUser         = "sinlov"
	gitHost         = "github.com"
)

// if not find config Path just try to use GOPATH code github.com/ShubNig/AubNig/config.conf
// if code aubnig.conf and run root path not found, return ""
func Try2FindOutConfigPath() (string, string, error) {
	configFilePath := filepath.Join(sCli.CommandPath(), defaultConfPath)
	projectPath := sCli.CurrentDirectory()
	if sFiles.IsFileExist(configFilePath) {
		return configFilePath, projectPath, nil
	}
	sCli.FmtYellow("\nWarning!\nCan not find config file at path: %s\n", sCli.CommandPath())
	goPathEnv := os.Getenv("GOPATH")
	goPathEnvS := strings.Split(goPathEnv, ":")
	isFindDevConf := false
	for _, path := range goPathEnvS {
		codePath := filepath.Join(path, "src", gitHost, gitUser, gitRepo)
		futurePath := filepath.Join(codePath, defaultConfPath)
		projectPath = filepath.Join(codePath, "build")
		if sFiles.IsFileExist(futurePath) {
			configFilePath = futurePath
			isFindDevConf = true
			break
		}
	}
	if isFindDevConf {
		sCli.FmtCyan("just use dev config at path: %s\n", configFilePath)
		return configFilePath, projectPath, nil
	} else {
		errInfo := fmt.Sprintf("can not load config at path: %s\nExit 1\n", configFilePath)
		configFilePath = ""
		return configFilePath, projectPath, errors.New(errInfo)
	}
}
