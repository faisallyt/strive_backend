package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic("Please set the environment variable " + key)
	}
	return val
}
func init() {
	fmt.Println("Loading envs...")
	LoadEnvs()
	if os.Getenv("ENV") == "production" {
		os.Setenv("GIN_MODE", "release")
		LoadEnvsFromAWS()
	} else if os.Getenv("ENV") == "development" {
		LoadEnvs()
	} else {
		log.Fatal("Please set the environment variable ENV")
	}

}

func LoadEnvs() (err error) {
	err = godotenv.Load()
	return
}
