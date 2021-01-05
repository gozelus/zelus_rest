package main

import (
	"fmt"
	"github.com/gozelus/zelus_rest/cli/actions"
	"github.com/urfave/cli"
	"os"
)

var apiCommand = cli.Command{
	Name:  "api",
	Usage: "根据 api 生成代码",
	Flags: []cli.Flag{
		cli.StringFlag{
			Required: true,
			Name:     "file",
			Usage:    `api 文件的入口`,
		},
		cli.StringFlag{
			Required: true,
			Name:     "appname",
			Usage:    `工程名`,
		},
	},
	Action: actions.GenApis,
}
var repoCommand = cli.Command{
	Name:  "repo",
	Usage: "生成数据库模型 repo 代码",
	Subcommands: []cli.Command{
		{
			Name:  "mysql",
			Usage: `根据 MySQL 表结构定义生成模型代码`,
			Subcommands: []cli.Command{
				{
					Name:  "datasource",
					Usage: `根据 MySQL 连接串链接`,
					Flags: []cli.Flag{
						cli.StringFlag{
							Required: true,
							Name:     "url",
							Usage:    `连接串, 如root:password@tcp(127.0.0.1:3306)/database`,
						},
						cli.StringFlag{
							Required: true,
							Name:     "table, t",
							Usage:    `表名`,
						},
					},
					Action: actions.GenRepo,
				},
			},
		},
	},
}

var modelCommand = cli.Command{
	Name:  "model",
	Usage: "生成数据库模型代码",
	Subcommands: []cli.Command{
		{
			Name:  "mysql",
			Usage: `根据 MySQL 表结构定义生成模型代码`,
			Subcommands: []cli.Command{
				{
					Name:  "datasource",
					Usage: `根据 MySQL 连接串链接`,
					Flags: []cli.Flag{
						cli.StringFlag{
							Required: true,
							Name:     "url",
							Usage:    `连接串, 如root:password@tcp(127.0.0.1:3306)/database`,
						},
						cli.StringFlag{
							Required: true,
							Name:     "table, t",
							Usage:    `表名`,
						},
					},
					Action: actions.GenModel,
				},
			},
		},
	},
}

var commands = []cli.Command{
	apiCommand,
	modelCommand,
	repoCommand,
}

func main() {
	app := cli.NewApp()
	app.Usage = "a cli tool to generate code"
	app.Commands = commands
	// cli already print error messages
	if err := app.Run(os.Args); err != nil {
		fmt.Println("error:", err)
	}
}
