package task

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
)

// CmdDone mark task as fnished
func CmdDone(c *cli.Context) {

	if len(c.Args()) != 1 {
		return
	}

	id := c.Args()[0]
	fmt.Printf("Task %s is done\n", id)

	db := dbConn()
	defer db.Close()

	_, err := db.Exec("UPDATE todos SET is_done = 1 WHERE id = " + id)
	if err != nil {
		log.Fatal(err)
	}
}
