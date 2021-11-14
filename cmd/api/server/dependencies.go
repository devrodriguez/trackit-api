package server

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Dependencies struct {
	sqlDB *gorm.DB
}

func BuildDependencies() Dependencies {
	db := buildMySQLConn()

	return Dependencies{
		db,
	}
}

func buildMySQLConn() *gorm.DB {
	// Build SQL Connection
	dsn := "root:@tcp(127.0.0.1:3306)/checkit?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
