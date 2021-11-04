package task

import (
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

var subCommands = []string{"all"}

// CmdList list all the tasks
func CmdList(c *cli.Context) {
	if len(c.Args()) != 0 {
		return
	}

	isAllMode, err := strconv.ParseBool(c.String("all"))
  checkError(err)

  rows := GetAll(isAllMode)
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
