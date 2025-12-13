package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"strings"
	"testing"
	"tui-app/app"
	"tui-app/storage/sqlite"
)

var (
	db      *sql.DB
	driver  string = "sqlite3"
	path    string = "storage/sqlite/test.db"
	err     error
	service app.TopicService
	store   app.TopicStore
)

func TestMain(m *testing.M) {
	db, err = sql.Open(driver, path)
	if err != nil {
		log.Fatalf("Couldn't connect to database: %s\n", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Datbase not reachable: %s\n", err)
	}
	store = sqlite.NewSQLiteTopicStorage(db)
	service = *app.NewTopicService(store)

	exitCode := m.Run()
	db.Close()
	os.Exit(exitCode)
}

func TestListTopics(t *testing.T) {
	var sb strings.Builder
	expected := `- calculus
- algebra
- databases
- tdd-ddd
`
	t.Run("Listing all Topics", func(t *testing.T) {
		service.ListTopics(&sb, context.Background())
		got := sb.String()
		if got != expected {
			t.Error("List of topic don't match")
		}
	})
}
