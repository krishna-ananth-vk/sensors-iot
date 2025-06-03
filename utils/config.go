package utils

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

var (
    Port      string
    DBPath    string
    AuthToken string
)

func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system env")
    }

    Port = os.Getenv("PORT")
    if Port == "" {
        Port = "8080"
    }

    DBPath = os.Getenv("DB_PATH")
    if DBPath == "" {
        DBPath = "./temperature.db"
    }

    AuthToken = os.Getenv("AUTH_TOKEN")
    if AuthToken == "" {
        log.Fatal("AUTH_TOKEN not set")
    }
}
