package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
	"zoro/pkg/config"
	"zoro/pkg/model"
)

type Handler struct {
	DB *gorm.DB
}

func InitConnectDB(config config.Config) Handler {
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.DBUserName,
		config.DBPassWord,
		config.DBHost,
		config.DBPort,
		config.DBName)
	db, err := gorm.Open(mysql.Open(connectString), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	log.Print("Connect Database success.")

	db.AutoMigrate(&model.Info{})
	db.AutoMigrate(&model.Purchase{})

	return Handler{db}
}
