package server

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Dependencies struct {
	sqlDB *sql.DB
}

func BuildDependencies() Dependencies {
	db := buildMySQLConn()

	return Dependencies{
		db,
	}
}

func buildMySQLConn() *sql.DB {
	// Build SQL Connection
	conf := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "checkit",
	}
	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatalln(pingErr)
	}

	return db
}
