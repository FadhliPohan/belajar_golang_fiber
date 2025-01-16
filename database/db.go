package database

import (
    "log"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)

func Connect() *gorm.DB {
    // Memuat variabel lingkungan dari file .env
    err := godotenv.Load("config/.env") // Menentukan path ke file .env
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Mengambil variabel lingkungan
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    // Membuat string koneksi
    dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

    // Menghubungkan ke database
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to database:", err)
    }
    return db
}
