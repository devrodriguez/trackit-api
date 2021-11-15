package server

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
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

	runMigrations(db)

	return db
}

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(&entity.CheckType{})
	db.AutoMigrate(&entity.Company{})
	db.AutoMigrate(&entity.Check{})
	db.AutoMigrate(&entity.Employee{})

}
