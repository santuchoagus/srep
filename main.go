package main

import (
	"context"
	"github.com/santuchoagus/srep/app"
	"github.com/santuchoagus/srep/storage/sqlite"
	"log"
	"os"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	devNull, _ := os.OpenFile("/dev/null", os.O_WRONLY, 0o666)

	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		log.Fatal("Error: HOME environment variable is not set.")
		return
	}

	dbPath := homeDir + "/" + ".local/share/srep/app.db"

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	store := sqlite.NewSQLiteTopicStorage(db)
	service := app.NewTopicService(store)
	service.ListTopicsVerbose(devNull, context.Background())

	app.StartCli(service)
}
