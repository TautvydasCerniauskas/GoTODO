package task

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

// CmdList list all the tasks
func CmdList(c *cli.Context) {
	if len(c.Args()) != 0 {
		return
	}

	isAllMode := c.String("all")
	var s string
	if isAllMode == "true" {
		s = "SELECT id, title, priority, is_done, created_at, updated_at FROM todos ORDER BY FIELD(priority, 'HIGH', 'MEDIUM', 'LOW')"
	} else {
		s = "SELECT id, title, priority, is_done, created_at, updated_at FROM todos WHERE is_done = 0 ORDER BY FIELD(priority, 'HIGH', 'MEDIUM', 'LOW')"
	}

	db := dbConn()
	defer db.Close()

	rows, err := db.Query(s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	data := [][]string{}

	for rows.Next() {
		var id, isDone int
		var title, priority string
		var createdAt, updatedAt int64
		rows.Scan(&id, &title, &priority, &isDone, &createdAt, &updatedAt)
		if createdAt == updatedAt {
			data = append(data, []string{strconv.Itoa(id), priority, title, doneLabel(isDone), dateForView(createdAt), ""})
		} else {
			data = append(data, []string{strconv.Itoa(id), priority, title, doneLabel(isDone), dateForView(createdAt), dateForView(updatedAt)})
		}
	}
	if len(data) > 0 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"No", "Priority", "Title", "Status", "Created", "Updated"})
		table.SetBorder(false)
		table.AppendBulk(data)
		table.Render()
	}
}

func doneLabel(isDone int) string {
	if isDone == 0 {
		return "-"
	}
	return "Done"
}

func dateForView(at int64) string {
	return time.Unix(at, 0).Format("2006-01-02 15:04:05")
}
