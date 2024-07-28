package main

import (
	"log"
	"net/http"

	"github.com/go_boilerplate/internal/config"
	"github.com/go_boilerplate/internal/db"
	"github.com/go_boilerplate/internal/discovery"
	"github.com/go_boilerplate/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	// Connect to the database
	var database db.DB = &db.GormDB{}
	gormDB, err := database.Connect(cfg.Database.DSN)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Setup service discovery using ServiceRegistry interface
	var registry discovery.ServiceRegistry
	registry, err = discovery.NewConsulServiceRegistry(cfg.Consul.Address)
	if err != nil {
		log.Fatalf("Could not create Consul service registry: %v", err)
	}

	// Register the service
	err = registry.Register(cfg.Consul.ServiceID, cfg.Consul.ServiceName, "localhost", 8080)
	if err != nil {
		log.Fatalf("Could not register service: %v", err)
	}

	// Setup Gin router
	r := router.SetupRouter(gormDB)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))

	// Deregister the service on shutdown
	defer registry.Deregister(cfg.Consul.ServiceID)
}
