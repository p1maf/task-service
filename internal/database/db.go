package database

import (
	"github.com/p1maf/task-service/internal/task"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		// Пример для локальной разработки
		dsn = "host=localhost user=user password=password dbname=usersdb port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB = db

	// Автоматическая миграция модели User
	if err := DB.AutoMigrate(&task.Task{}); err != nil {
		log.Fatalf("failed to migrate Task model: %v", err)
	}
}
