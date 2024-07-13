package main

import (
	"log"
	"os"
	"go-api-gdrive/utils"
	"github.com/joho/godotenv"
)


func main() {

	err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    fileID := os.Getenv("ID")

    if err := utils.DownloadFile(fileID); err != nil {
        log.Fatalf("Unable to download file: %v", err)
    }
}

