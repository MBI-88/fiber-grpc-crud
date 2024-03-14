package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
	fileLog *os.File
	loggerError = log.New(fileLog, "Database Error =>", log.Ldate)
)

func setConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsn),&gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic(err)
	}

	return db
}

func createLogs(path string) {
	fileLog, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	loggerError.SetOutput(fileLog)
}

func Migrate(url string) {
	DB = setConnection(url)
	if err := DB.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	fmt.Println("[+] Migration done!")
}

func DialDb(url, logpath string) {
	DB = setConnection(url)
	go createLogs(logpath)
}