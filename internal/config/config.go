package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load the env for the storage path
func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Env can't be loaded")
	}
	storagePath := os.Getenv("STORAGE_PATH")
	if storagePath == "" {
		// storage path is empty
		log.Fatal("Storage path is empty")
	}

	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		log.Fatal("Storage file not found in set storage path")
	}

	fmt.Println("Env loaded")
	fmt.Println(os.Getenv("STORAGE_PATH"))
}
