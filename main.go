package main

import (
	"github/luqxus/spxce/api"
	"github/luqxus/spxce/database"
	"github/luqxus/spxce/service"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {

	// load environment
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// database configuration
	datastoreConfig := newDatastoreConfig()

	// new Postgres database store
	datastore, err := database.NewPGDatabase(datastoreConfig)
	if err != nil {
		log.Fatal(err)
	}

	// new service
	service := service.New(datastore)

	// new api server
	api := api.New(api.APIServerConfig{
		Port:    3000,
		Host:    "192.168.1.107",
		Service: service,
	})

	// start api server
	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}

// parse and return database configs from env
func newDatastoreConfig() database.DatabaseConfig {

	host := os.Getenv("DATABASE_HOST")
	if host == "" {
		log.Fatal("DATABASE_HOST not an environment variable")
	}

	port := os.Getenv("DATABASE_PORT")
	if host == "" {
		log.Fatal("DATABASE_PORT not an environment variable")
	}

	user := os.Getenv("DATABASE_USER")
	if host == "" {
		log.Fatal("DATABASE_USER not an environment variable")
	}

	password := os.Getenv("DATABASE_PASSWORD")
	if host == "" {
		log.Fatal("DATABASE_PASSWORD not an environment variable")
	}

	databaseName := os.Getenv("DATABASE_NAME")
	if host == "" {
		log.Fatal("DATABASE_NAME not an environment variable")
	}

	p, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal(err)
	}

	return database.DatabaseConfig{
		Host:     host,
		Port:     p,
		User:     user,
		Password: password,
		DBName:   databaseName,
	}
}
