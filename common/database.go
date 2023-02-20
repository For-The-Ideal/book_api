package common

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "test"
	username := "root"
	password := "123456"
	charset := "utf8"
	local := "Local"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s", username, password, host, port, database, charset, local)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	sqlDb, _ := db.DB()
	if err != nil {
		panic("open mysql failed," + err.Error())
	}
	pingErr := sqlDb.Ping()
	if pingErr != nil {
		panic("mysql link error," + err.Error())
	}
	sqlDb.SetMaxOpenConns(5)
	sqlDb.SetMaxIdleConns(2)
	sqlDb.SetConnMaxIdleTime(time.Minute)
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
