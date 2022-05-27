package postgresstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			exec, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
			if err != nil {
				fmt.Println("error clean test db", exec, err)
			}
		}
		db.Close()
	}
}
