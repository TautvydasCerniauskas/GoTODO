package task

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
)

// CmdAdd add new task
func CmdAdd(c *cli.Context) {
	if len(c.Args()) == 0 {
		fmt.Println("[ERROR] Must set task title")
		return
	}

	title := strings.Join(c.Args(), " ")
	db := dbConn()
	defer db.Close()

	now := strconv.FormatInt(time.Now().Unix(), 10)
	tag := getTag(title)
  title = trimedTitle(title)
	fmt.Println("Added a new task: ", title)

	_, err := db.Exec("INSERT INTO todos(title, priority, is_done, created_at, updated_at) values('" + title + "','" + tag + "', 0, " + now + ", " + now + ")")
  checkError(err)
}
