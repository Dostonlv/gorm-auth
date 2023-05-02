package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	ServiceName string
	ServiceHost string
	ServicePort string

	SecretKey string

	Environment string
	Version     string

	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	PostgresPort     string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("ErrEnvNodFound")
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "gorm-auth"))
	config.ServiceHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "localhost"))
	config.ServicePort = cast.ToString(getOrReturnDefaultValue("SERVICE_PORT", ":8080"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "DebugMode"))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.SecretKey = cast.ToString(getOrReturnDefaultValue("SECRET", "secret"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "0.0.0.0"))
	config.PostgresPort = cast.ToString(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "postgres"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", config.ServiceName))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
