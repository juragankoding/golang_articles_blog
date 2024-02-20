package env

import (
	"os"
)

func GetMongoDBURL() string {
	return os.Getenv("MONGO_URL")
}

func GetMongoDatabase() string {
	return os.Getenv("MONGO_DATABASE")
}
