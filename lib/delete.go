package task

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
)

// CmdDelete list all tasks
func CmdDelete(c *cli.Context) {

	isAllMode := c.String("all")

	var id string
	if len(c.Args()) == 1 {
		id = c.Args()[0]
	}

	var s, m string
	if isAllMode == "true" {
		s = "TRUNCATE TABLE todos"
		m = "Deleted all tasks"
	} else {
		s = "DELETE FROM todos WHERE ID = " + id
		m = "Deleted task with id " + id
	}

	db := dbConn()
	defer db.Close()

	_, err := db.Exec(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}
