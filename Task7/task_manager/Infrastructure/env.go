package Infrastructure

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"strconv"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	// ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	MongoAPIURL            string `mapstructure:"MONGO_API_URL"`
	// DBHost                 string `mapstructure:"DB_HOST"`
	// DBPort                 string `mapstructure:"DB_PORT"`
	// DBUser                 string `mapstructure:"DB_USER"`
	// DBPass                 string `mapstructure:"DB_PASS"`
	// DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	// RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	// RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() *Env {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found or error loading it")
	}
	contextTimeout, err := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	if err != nil {
		log.Printf("Error converting CONTEXT_TIMEOUT to int: %v", err)
		contextTimeout = 0
	}
	accessTokenExpiryHour, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"))
	if err != nil {
		log.Printf("Error converting ACCESS_TOKEN_EXPIRY_HOUR to int: %v", err)
		accessTokenExpiryHour = 0
	}
	env := &Env{
		AppEnv:                os.Getenv("APP_ENV"),
		ContextTimeout:        contextTimeout,
		MongoAPIURL:           os.Getenv("MONGO_API_URL"),
		AccessTokenExpiryHour: accessTokenExpiryHour,
		AccessTokenSecret:     os.Getenv("ACCESS_TOKEN_SECRET"),
	}
	

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return env
}