package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/mengdj/goctl-rest-discover/generate"
	"github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
)

var (
	version = "v0.0.1"
	author  = []*cli.Author{
		&cli.Author{
			Name:  "mengdj",
			Email: "mengdj@outlook.com",
		},
	}
	commands = []*cli.Command{
		{
			Name:  "rest-discover",
			Usage: "generates rest discover client factory",
			Action: func(context *cli.Context) error {
				plugin, err := plugin.NewPlugin()
				if nil != err {
					return err
				}
				return generate.Do(plugin)
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "package",
					Usage: "the package of rest discover",
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Authors = author
	app.Usage = "a plugin of goctl to generate rest discover client"
	app.Version = fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("goctl-rest-discover: %+v\n", err)
	}
}
