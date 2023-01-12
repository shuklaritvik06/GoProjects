package commands

import (
	"github.com/shuklaritvik06/GoProjects/GO_CLI/pkg/config"
	"github.com/shuklaritvik06/GoProjects/GO_CLI/pkg/scraper"
	"github.com/urfave/cli"
)

func Commands() {
	app := config.GetApp()
	app.Commands = []cli.Command{
		{
			Name:    "getfollowers",
			Aliases: []string{"gfr"},
			Usage:   "Get List of the user who follow the GitHub Acount",
			Action: func(c *cli.Context) {
				scraper.InitScraper(c.Args().First(), "followers")
			},
		},
		{
			Name:    "getfollowing",
			Aliases: []string{"gfg"},
			Usage:   "Get List of the users that the given account follow",
			Action: func(c *cli.Context) {
				scraper.InitScraper(c.Args().First(), "following")
			},
		},
	}
}
