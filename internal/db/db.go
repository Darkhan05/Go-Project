package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"crudproject/internal/models"
)

var DB *gorm.DB

func InitDB() {
	dbHost := "localhost"
	dbName := "postgres"
	dbUser := "postgres"
	dbPass := "7982"
	dbPort := "5432"
	sslmode := "disable"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, sslmode)

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, sslmode)

	sqlDB, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("SQL Open error:", err)
	}

	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal("migratepg.WithInstance error:", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	migrationsPath := filepath.Join(cwd, "internal", "db", "migrations")
	migrationsPath = filepath.ToSlash(migrationsPath)
	migrationsURL := fmt.Sprintf("file://%s", strings.TrimPrefix(migrationsPath, "/"))

	m, err := migrate.NewWithDatabaseInstance(migrationsURL, "postgres", driver)
	if err != nil {
		log.Fatal("migrate.NewWithDatabaseInstance error:", err)
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("Migration up error:", err)
	}

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("gorm.Open error:", err)
	}

	err = gormDB.AutoMigrate(&models.Car{})
	if err != nil {
		log.Fatal("AutoMigrate error:", err)
	}

	DB = gormDB
}
