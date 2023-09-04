package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dbSource = "root:root@tcp(localhost:3306)/UserService?parseTime=true"
)

var testQueries *Queries
var testDB *sql.DB

func CreateTestDB() (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestMain(m *testing.M) {
	var err error
	testDB, err = CreateTestDB()

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
