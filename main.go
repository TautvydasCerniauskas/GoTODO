package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Email = ""
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
