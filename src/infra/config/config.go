package config

import (
	"github.com/joho/godotenv"

	"os"

	"fmt"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		fmt.Print("Error loading config file!")
	}
}

// Get func to get env value from key ---
func Get(key string) string {
	return os.Getenv(key)
}
