package main

import (
	"fmt"
	"github.com/gozelus/zelus_rest/cli/actions"
	"github.com/urfave/cli"
	"os"
)

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
						cli.StringFlag{
							Required: true,
							Name:     "pkg",
							Usage:    "package name",
						},
						cli.StringFlag{
							Required: true,
							Name:     "dir, d",
							Usage:    "目标文件夹",
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
						cli.StringFlag{
							Required: true,
							Name:     "pkg",
							Usage:    "package name",
						},
						cli.StringFlag{
							Required: true,
							Name:     "dir, d",
							Usage:    "目标文件夹",
						},
					},
					Action: actions.GenModel,
				},
			},
		},
	},
}

var commands = []cli.Command{
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
