package Database

import (
	"fmt"
	"mario/src/Entities/User"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ---------------------------------------------

var ORM *gorm.DB

// ---------------------------------------------

type DbConfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
}

// ---------------------------------------------

var dbConfig = DbConfig{
	Host:     "127.0.0.1",
	Port:     3306,
	DBName:   "go_mario",
	User:     "root",
	Password: "MARIOmario_123",
}

// ---------------------------------------------

var DbURL = fmt.Sprintf(
	"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
	dbConfig.User,
	dbConfig.Password,
	dbConfig.Host,
	dbConfig.Port,
	dbConfig.DBName,
)

// ---------------------------------------------

func DbConnectAndMigrate() {
	db, err := gorm.Open(mysql.Open(DbURL), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User.User{})

	ORM = db
}

// ---------------------------------------------
