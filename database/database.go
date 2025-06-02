package database

import (
    "fmt"
    "log"

    "github.com/chuks/BOTGO/config"
    "github.com/chuks/BOTGO/models"
    _ "github.com/lib/pq" // Import the PostgreSQL driver
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// Init initializes the database connection and returns a gorm.DB instance.
func Init() *gorm.DB {
    // Load configuration values
    host, err := config.LoadConfig("POSTGRES_HOST")
    if err != nil {
        log.Fatalf("failed to load POSTGRES_HOST: %v", err)
    }

    user, err := config.LoadConfig("POSTGRES_USER")
    if err != nil {
        log.Fatalf("failed to load POSTGRES_USER: %v", err)
    }

    password, err := config.LoadConfig("POSTGRES_PASSWORD")
    if err != nil {
        log.Fatalf("failed to load POSTGRES_PASSWORD: %v", err)
    }

    port, err := config.LoadConfig("POSTGRES_PORT")
    if err != nil {
        log.Fatalf("failed to load POSTGRES_PORT: %v", err)
    }

    // Connect to the initial database
    dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s",
        host, user, password, port, "postgres", // Default database for initial connection
    )

    DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    // Drop and recreate the database
    if err := DB.Exec("DROP DATABASE IF EXISTS todolist;").Error; err != nil {
        log.Fatalf("failed to drop database: %v", err)
    }

    if err := DB.Exec("CREATE DATABASE todolist;").Error; err != nil {
        log.Fatalf("failed to create database: %v", err)
    }

    // Update DSN for the newly created database
    dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
        host, user, password, "todolist", port, // Use the newly created database
    )

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to the new database: %v", err)
    }

    // Migrate tables
    if err := DB.AutoMigrate(&models.Task{}); err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }

    return DB
}