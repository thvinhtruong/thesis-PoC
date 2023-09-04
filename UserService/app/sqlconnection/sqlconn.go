package sqlconnection

import (
	"database/sql"
	"log"
	"server/UserService/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DBConn *sql.DB

func init() {
	var err error
	configuration := config.Singleton
	DBConn, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/UserService?parseTime=true")

	if err != nil {
		log.Fatal("cannot connect to db:", err)
		panic(err)
	}

	DBConn.SetMaxOpenConns(100)
	DBConn.SetMaxIdleConns(10)
	DBConn.SetConnMaxLifetime(time.Minute * 3)

	if err = DBConn.Ping(); err != nil {
		log.Fatal("cannot ping db:", err)
	}

	log.Printf("Connected to database: %v", configuration.GetConfig(config.MYSQL_DATABASE))
}

func GetDB() *sql.DB {
	return DBConn
}

func CloseDB() error {
	return DBConn.Close()
}

func PrintInfo(config config.SConfig, err error) {
	log.Println("Exit with error: ", err)
	log.Println("Reconnecting to database with info: ", config)
	time.Sleep(time.Second * 3)
}
