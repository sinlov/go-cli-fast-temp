package aubnig

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"os"
	"strings"
	"path/filepath"
)

func TestTry2FindOutConfigPath(t *testing.T) {
	convey.Convey("mock TestTry2FindOutConfigPath", t, func() {
		// mock
		goPathEnv := os.Getenv("GOPATH")
		goPathEnvS := strings.Split(goPathEnv, ":")
		goFirstPath := goPathEnvS[0]
		data := struct {
			wantConfigPath  string
			wantProjectFile string
		}{
			wantConfigPath:  filepath.Join(goFirstPath, "src", gitHost, gitUser, gitRepo, defaultConfPath),
			wantProjectFile: filepath.Join(goFirstPath, "src", gitHost, gitUser, gitRepo, "build"),
		}
		customPath := defaultConfPath
		customData := struct {
			custom          string
			wantConfigPath  string
			wantProjectFile string
		}{
			custom:          customPath,
			wantConfigPath:  filepath.Join(goFirstPath, "src", gitHost, gitUser, gitRepo, customPath),
			wantProjectFile: filepath.Join(goFirstPath, "src", gitHost, gitUser, gitRepo, "build"),
		}
		convey.Convey("do TestTry2FindOutConfigPath", func() {
			// do
			configPath, projectFile, err := Try2FindOutConfigPath("")
			customConfigPath, customProjectFile, errCustom := Try2FindOutConfigPath(customData.custom)
			convey.Convey("verify TestTry2FindOutConfigPath", func() {
				// verify
				if err != nil || errCustom != nil {
					t.Errorf("find config Error, %s, %s", err, errCustom)
				} else {
					convey.So(configPath, convey.ShouldEqual, data.wantConfigPath)
					convey.So(customConfigPath, convey.ShouldEqual, customData.wantConfigPath)
					convey.So(projectFile, convey.ShouldEqual, data.wantProjectFile)
					convey.So(customProjectFile, convey.ShouldEqual, customData.wantProjectFile)
				}

			})
		})
	})
}
