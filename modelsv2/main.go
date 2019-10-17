package modelsv2

import (
	"database/sql"
	"log"

	// import sql
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDb for start a pool connection
func InitDb(user string, password string, host string, database string) {
	var err error
	db, err = sql.Open("mysql", user+":"+password+"@tcp("+host+":3306)/"+database+"?parseTime=true")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	db.SetConnMaxLifetime(0)
}

// Close the pôol connection
func Close() {
	db.Close()
}
