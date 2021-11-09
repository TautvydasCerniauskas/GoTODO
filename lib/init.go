package task

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

// CmdInit create database and tables.
func CmdInit(c *cli.Context) {

	db := dbConn()
	defer db.Close()

	createTable := `
	CREATE TABLE todos (
		id integer PRIMARY KEY AUTOINCREMENT,
		title text NOT NULL,
		priority text NOT NULL,
		is_done integer NOT NULL,
		created_at integer NOT NULL,
		updated_at integer NOT NULL
	);
	`
	_, err := db.Exec(createTable)
	if err != nil {
		log.Printf("%q: %s\n", err, createTable)
		return
	}
}

func dbConn() (db *sql.DB) {
	err := godotenv.Load()
	checkError(err)
	db, err = sql.Open("sqlite3", "./todo.db")
	checkError(err)
	return db
}
