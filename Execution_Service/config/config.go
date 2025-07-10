package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	DUMMY_API_URL                string
	OTHENTIC_CLIENT_RPC_ADDRESS string
	PrivateKey                  string
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	DUMMY_API_URL = os.Getenv("DUMMY_API_URL")
	OTHENTIC_CLIENT_RPC_ADDRESS = os.Getenv("OTHENTIC_CLIENT_RPC_ADDRESS")
	PrivateKey = os.Getenv("PRIVATE_KEY_PERFORMER")

	if DUMMY_API_URL == "" || OTHENTIC_CLIENT_RPC_ADDRESS == "" || PrivateKey == "" {
		log.Fatal("Environment variables are not set properly")
	}

	gin.SetMode(gin.ReleaseMode)
}