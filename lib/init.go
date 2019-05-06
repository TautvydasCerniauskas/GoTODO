package task

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

// CmdInit create database and tables.
func CmdInit(c *cli.Context) {

	db := dbConn()
	defer db.Close()

	createTable := `
	CREATE TABLE todos (
		id integer AUTO_INCREMENT PRIMARY KEY,
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
	if err != nil {
		log.Fatal("Error with a .env file")
	}
	dbDriver := os.Getenv("DBDRIVER")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASSWORD")
	dbName := os.Getenv("DBNAME")
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
