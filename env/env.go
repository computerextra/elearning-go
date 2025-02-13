package env

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URL  string `env:"DATABASE_URL,required"`
	PORT          int    `env:"PORT,required"`
	MAIL_FROM     string `env:"MAIL_FROM,required"`
	MAIL_SERVER   string `env:"MAIL_SERVER,required"`
	MAIL_PORT     int    `env:"MAIL_PORT,required"`
	MAIL_USER     string `env:"MAIL_USER,required"`
	MAIL_PASSWORD string `env:"MAIL_PASSWORD,required"`
}

func GetEnv() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	cfg := Config{} // 👈 new instance of `Config`

	err = env.Parse(&cfg) // 👈 Parse environment variables into `Config`
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	return cfg
}
