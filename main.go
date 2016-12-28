package main

import (
	"github.com/mkideal/cli"
	clix "github.com/mkideal/cli/ext"
	"os"
	"fmt"
)

const (
	versionName string = "1.0.0"
	//TODO set comm info
	commInfo string = "XXX utils"
)

type config struct {
	A string
	B int
	C bool
}

type filterCLI struct {
	cli.Helper
	Version bool `cli:"version" usage:"version"`
	Verbose bool `cli:"verbose" usage:"see Verbose of utils"`
	Port int  `cli:"p,port" usage:"short and long format flags both are supported"`
	Id uint8 `cli:"*id" usage:"this is a required flag, note the *"`
	Env    string `cli:"env" usage:"env variable as default" dft:"$HOME"`
	Expr   int    `cli:"expr" usage:"expression as default" dft:"$BASE_PORT+1000"`
	DevDir string `cli:"devdir" usage:"directory of developer" dft:"$HOME/dev"`
	Force  bool `cli:"!v" usage:"force flag, note the !"`
	Required int  `cli:"*r" usage:"required flag"`

	Username string `cli:"u,username" usage:"github account" prompt:"type github account"`
	Password string `pw:"password" usage:"password of github account" prompt:"type the password"`

	PidFile clix.PidFile `cli:"pid" usage:"pid file" dft:"013-pidfile.pid"`
	Time     clix.Time     `cli:"t" usage:"time"`
	Duration clix.Duration `cli:"d" usage:"duration"`

	Content clix.File `cli:"f,file" usage:"read content from file or stdin"`

	JSON config `cli:"c,config" usage:"parse json string" parser:"json"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Warning you input is error pleae use -h to see help")
		os.Exit(-1)
	}
	cli.Run(new(filterCLI), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*filterCLI)
		if argv.Version {
			ctx.String(commInfo + "\n\tversion: " + versionName)
			os.Exit(0)
		}
		return nil
	})
}
