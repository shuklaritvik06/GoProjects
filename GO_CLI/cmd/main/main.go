package main

import (
	"log"
	"os"

	"github.com/shuklaritvik06/GoProjects/GO_CLI/pkg/commands"
	"github.com/shuklaritvik06/GoProjects/GO_CLI/pkg/config"
)

func main() {
	config.Info()
	commands.Commands()
	app := config.GetApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
