package task

import (
	"fmt"
	"log"
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
	dbTitle := title[1:]
	fmt.Println("Added a new task: ", dbTitle)

	_, err := db.Exec("INSERT INTO todos(title, priority, is_done, created_at, updated_at) values('" + dbTitle + "','" + tag + "', 0, " + now + ", " + now + ")")
	if err != nil {
		log.Fatal(err)
	}
}

func getTag(title string) string {
	if strings.Contains(title, "!") {
		return HIGH.String()
	} else if strings.Contains(title, "$") {
		return LOW.String()
	}
	return MEDIUM.String()
}
