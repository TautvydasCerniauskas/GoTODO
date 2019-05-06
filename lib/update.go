package task

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
)

// CmdUpdate will update a specific task
func CmdUpdate(c *cli.Context) {
	if len(c.Args()) < 2 {
		fmt.Println("[ERROR] Must set task id and title")
	}

	id := c.Args()[0]
	title := strings.Join(c.Args()[1:], " ")

	db := dbConn()

	defer db.Close()

	now := strconv.FormatInt(time.Now().Unix(), 10)

	_, err := db.Exec("UPDATE todos SET title = '" + title + "', updated_at = " + now + " WHERE id = " + id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated task id = ", id)
}
