package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type LamodaConfig struct {
	HTTPServer *HttpServer
	DB         *Db
}

type HttpServer struct {
	Host string `envconfig:"HOSTNAME" default:"127.0.0.1"`
	Port uint16 `envconfig:"PORT" default:"8080"`
}

type Db struct {
	DBUser            string        `envconfig:"POSTGRES_USER" required:"true"`
	DBPass            string        `envconfig:"POSTGRES_PASS" required:"true"`
	DBName            string        `envconfig:"POSTGRES_NAME" required:"true"`
	DBPort            string        `envconfig:"POSTGRES_PORT" required:"true"`
	DBHost            string        `envconfig:"POSTGRES_HOST" required:"true"`
	DBMaxOpenConns    int           `envconfig:"POSTGRES_MAX_OPEN_CONNS" default:"10"`
	DBMaxIdleConns    int           `envconfig:"POSTGRES_MAX_IDLE_CONNS" default:"2"`
	DBConnMaxLifetime time.Duration `envconfig:"POSTGRES_CONN_MAX_LIFETIME" default:"60s"`
}

func getEnvFilenames() []string {
	return []string{".env.local", ".env"}
}

func (myDb *Db) GetDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		myDb.DBUser, myDb.DBPass, myDb.DBHost, myDb.DBPort, myDb.DBName,
	)
}

func LoadConfig(ctx context.Context) (*LamodaConfig, error) {
	for _, fileName := range getEnvFilenames() {
		err := godotenv.Load(fileName)
		if err != nil {
			log.Printf("error loading %s fileName : %v", fileName, err)
		}
	}

	var cfg LamodaConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("cannot process envs: %v", err)

		return nil, fmt.Errorf("cannot process envs: %w", err)
	} else {
		log.Printf("Config initialized")
	}

	return &cfg, nil
}
