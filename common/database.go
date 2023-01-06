package common

import (
	gomysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDatabase() {
	config := gomysql.Config{
		User:                 GetDotEnvValue("DBUSER"),
		Passwd:               GetDotEnvValue("DBPASS"),
		Net:                  GetDotEnvValue("DBNET"),
		Addr:                 GetDotEnvValue("DBADDRESS"),
		DBName:               GetDotEnvValue("DBNAME"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: config.FormatDSN(),
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *gorm.DB {
	return DB
}
