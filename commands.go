package main

import (
	"fmt"
	"os"

	lib "github.com/tautvydascerniauskas/TODO/lib"
	"github.com/urfave/cli"
)

// GlobalFlags is global option for all command
var GlobalFlags = []cli.Flag{}

var forceFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "force, f",
		Usage: "Force initialize even if database already exists",
	},
}

var allListFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "all, a",
		Usage: "List all tasks.",
	},
}

var allDeteleFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "all, a",
		Usage: "Delete all flags",
	},
}

// Commands are slice of command
var Commands = []cli.Command{
	{
		Name:   "init",
		Usage:  "Initialize SQLITE Database",
		Action: lib.CmdInit,
		Flags:  forceFlags,
	},
	{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "Add task",
		Action:  lib.CmdAdd,
		Flags:   []cli.Flag{},
	},
	{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "List all tasks",
		Action:  lib.CmdList,
		Flags:   allListFlags,
	},
	{
		Name:    "done",
		Aliases: []string{"d"},
		Usage:   "Mark a task as done",
		Action:  lib.CmdDone,
		Flags:   []cli.Flag{},
	},
	{
		Name:   "delete",
		Usage:  "Delete a task",
		Action: lib.CmdDelete,
		Flags:  allDeteleFlags,
	},
	{
		Name:    "update",
		Aliases: []string{"u"},
		Usage:   "Update a specific task",
		Action:  lib.CmdUpdate,
		Flags:   []cli.Flag{},
	},
}

// CommandNotFound is custom error
func CommandNotFound(c *cli.Context, command string) {
  name := c.App.Name
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. see '%s --help'.", name, command, name, name)
	os.Exit(2)
}
