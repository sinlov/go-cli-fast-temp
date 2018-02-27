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
			wantConfigPath:  filepath.Join(goFirstPath, "src", "github.com", "sinlov", "go-cli-fast-temp", "temp.conf"),
			wantProjectFile: filepath.Join(goFirstPath, "src", "github.com", "sinlov", "go-cli-fast-temp", "build"),
		}
		convey.Convey("do TestTry2FindOutConfigPath", func() {
			// do
			configPath, projectFile, err := Try2FindOutConfigPath()
			convey.Convey("verify TestTry2FindOutConfigPath", func() {
				// verify
				if err != nil {
					t.Errorf("find config Error, %s", err)
				} else {
					convey.So(configPath, convey.ShouldEqual, data.wantConfigPath)
					convey.So(projectFile, convey.ShouldEqual, data.wantProjectFile)
				}

			})
		})
	})
}
