package main

import (
	"github.com/docker/go-plugins-helpers/network"
	"github.com/urfave/cli"
	"github.com/joely1101/docker-plugin-hostnic/driver"
	"github.com/joely1101/docker-plugin-hostnic/log"
	"os"
)

const (
	version = "0.1"
)

func main() {

	var flagDebug = cli.BoolFlag{
		Name:  "debug, d",
		Usage: "enable debugging",
	}
	app := cli.NewApp()
	app.Name = "hostnic"
	app.Usage = "Docker Host Nic Network Plugin"
	app.Version = version
	app.Flags = []cli.Flag{
		flagDebug,
	}
	app.Action = Run
	app.Run(os.Args)
}

// Run initializes the driver
func Run(ctx *cli.Context) {
	if ctx.Bool("debug") {
		log.SetLevel("debug")
	}
	log.Info("Run %s", ctx.App.Name)
	d, err := driver.New()
	if err == nil {
		h := network.NewHandler(d)
		err = h.ServeUnix("root", 0)
	}
	if err != nil {
		log.Fatal("Run app error: %s", err.Error())
		os.Exit(1)
	}
}
