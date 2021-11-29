package main

import (
	"os"
	"runtime"

	"github.com/gabehoban/go-cli-template/cmd"
	"github.com/gabehoban/go-cli-template/core"
	"github.com/urfave/cli/v2"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.App{
		Name:                 core.AppConfig.GetString("app", "name"),
		Usage:                core.AppConfig.GetString("app", "desc"),
		Version:              core.AppConfig.GetString("app", "version"),
		Author:               core.AppConfig.GetString("app", "author"),
		Email:                core.AppConfig.GetString("app", "email"),
		EnableBashCompletion: true,
		Commands: []cli.Command{
			cmd.Hello,
		},
	}

	app.Run(os.Args)
}
