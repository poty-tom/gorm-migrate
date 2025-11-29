package main

import (
	"fmt"
	"log"
	"migrate/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	user := os.Getenv("MARIADB_USER")
	password := os.Getenv("MARIADB_PASSWORD")
	host := os.Getenv("MARIADB_HOST")
	port := os.Getenv("MARIADB_PORT")
	database := os.Getenv("MARIADB_DATABASE")

	// 接続
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	log.Println("Database connection established")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get database instance: ", err)
	}
	defer sqlDB.Close()

	// マイグレーション実行
	log.Println("Starting migration...")
	if err := db.AutoMigrate(
		&model.User{},
	); err != nil {
		log.Fatal("Migration failed: ", err)
	}
	log.Println("Migration completed successfully")
}
