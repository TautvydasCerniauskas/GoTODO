package task

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/urfave/cli"
)

// CmdUpdate will update a specific task
func CmdUpdate(c *cli.Context) {
	if len(c.Args()) < 2 {
    err := errors.New("Must set task id and title")
    log.Fatal().
        Err(err).
        Msg("service")
	}

	id := c.Args()[0]
	title := strings.Join(c.Args()[1:], " ")

	todo := TodoUpdate{
		id:    id,
		title: title,
	}
	UpdateTodo(todo)
	fmt.Println("Updated task id = ", id)
}
