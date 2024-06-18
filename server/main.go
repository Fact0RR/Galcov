package main

import (
	"log"
	"os"
	"server/handlers"
	"server/server"
	"server/storage"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v2"
)


func main() {
	filename := "./conf/conf.yaml"
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	// Распаковка данных YAML в структуру
	var config server.MainConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error unmarshaling data: %v", err)
	}

	server:=server.Server{
		App: fiber.New(),
		Config: &config,
		Store: storage.New(config.DataBaseString),
	}

	server.Store.Open()
	
	handler := handlers.Handler{Server: &server}

	handler.Routers()

	server.App.Listen(server.Config.Port)
}

