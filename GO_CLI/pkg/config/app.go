package config

import (
	"fmt"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func Info() {
	app.Name = "GitHub Scraper"
	app.Usage = "A CLI Scraper for GitHub"
	app.Author = "shuklaritvik"
	app.Version = "1.0.0"
	app.Action = func(*cli.Context) error {
		fmt.Print("\nINFO:// Please give the required arguements\n")
		return nil
	}
	app.Before = func(ctx *cli.Context) error {
		fmt.Print("Cooking....\n")
		return nil
	}
}

func GetApp() *cli.App {
	return app
}
