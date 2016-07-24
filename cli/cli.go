package cli

import (
	"github.com/codegangsta/cli"
	"github.com/Sirupsen/logrus"
	"path"
	"os"
)

func Run() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Usage = "dockerDemo"
	app.Version = "dockerDemo v0"
	app.Author = "wangforthinker"
	app.Email = "wangforthinker@163.com"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name : "log-level",
			Value: "info",
			EnvVar: "DEMO_API_LOG_LEVEL",
			Usage: "Log level (options: debug, info, warn, error, fatal, panic),",
		},
	}

	app.Before = func(c *cli.Context) error {
		logrus.SetOutput(os.Stderr)
		level,err := logrus.ParseLevel(c.String("log-level"))
		if err != nil {
			logrus.Fatal(err.Error())
		}

		logrus.SetLevel(level)

		return nil
	}

	app.Commands = []cli.Command {
		cli.Command{
			Name : "start",
			Action : startCommand,
			Flags: []cli.Flag {
				flDockerScheme,
				flDockerEndpoint,
				flDockerTls,
				flDockerTlsCertFile,
				flDockerTlsKeyFile,
			},

		},

	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}