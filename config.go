package sync

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment       string // develop, staging, production
	ServiceName       string
	OrderServiceUrl   string
	FinanceServiceUrl string
	PaymeServiceUrl   string

	LogLevel string

	HttpPort string

	OfdUrl string
}

// Load loads environment vars and inflates Config
func Load() Config {
	envFileName := cast.ToString(getOrReturnDefault("ENV_FILE_PATH", "../../.env"))

	if err := godotenv.Load(envFileName); err != nil {
		fmt.Println("No .env file found")
	}
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.ServiceName = cast.ToString(getOrReturnDefault("SERVICE_NAME", "billz_sync_service"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8090"))

	c.OfdUrl = cast.ToString(getOrReturnDefault("OFD_ENDPOINT", "http://127.0.0.1:3448/rpc/api"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
