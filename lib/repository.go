package task

import (
	"database/sql"
	"strconv"
	"time"
)

type Todo struct {
	id         int
	title      string
	priority   string
	done       int
	created_at int64
	updated_at int64
}

type TodoUpdate struct {
	title string
	id    string
	done  bool
}

func GetById(id string) Todo {
  var isDone int
  var title, priority string
  var createdAt, updatedAt int64

	db := dbConn()
	defer db.Close()

	rows := db.QueryRow("SELECT * FROM todos WHERE id = $1", id)
  rows.Scan(&id, &title, &priority, &isDone, &createdAt, &updatedAt)
  newId, err := strconv.Atoi(id)
  checkError(err)
  todo := Todo{
    id: newId,
    title: title,
    priority: priority,
    done: isDone,
    created_at: createdAt,
    updated_at: updatedAt,
  }

	return todo
}

func GetAll(all bool) *sql.Rows {
	var s string
	if all == true {
		s = "SELECT id, title, priority, is_done, created_at, updated_at FROM todos ORDER BY priority"
	} else {
		s = "SELECT id, title, priority, is_done, created_at, updated_at FROM todos WHERE is_done = 0 ORDER BY priority"
	}
	db := dbConn()
	defer db.Close()

	rows, err := db.Query(s)
	checkError(err)

	return rows
}

func UpdateTodo(t TodoUpdate) {
	db := dbConn()

	defer db.Close()

	now := strconv.FormatInt(time.Now().Unix(), 10)

	_, err := db.Exec("UPDATE todos SET title = '" + t.title + "', updated_at = " + now + " WHERE id = " + t.id)
	checkError(err)
}
