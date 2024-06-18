package server

import (
	"server/storage"

	"github.com/gofiber/fiber/v2"
)

type MainConfig struct {
	Port           string `yaml:"port"`
	DataBaseString string `yaml:"dataBaseString"`
}

type Server struct {
	App    *fiber.App
	Config *MainConfig
	Store  *storage.Store
}

