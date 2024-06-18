package handlers

import (
	"server/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Routers() {
	h.Server.App.Post("/auth", h.Auth)
	h.Server.App.Post("/reg",h.Reg)
	h.Server.App.Get("/", func(c *fiber.Ctx) error { return c.SendString("ok") })
	h.Server.App.Get("/prc",h.Price)
	h.Server.App.Use(middleware.LoginMiddleware)
	h.Server.App.Get("/hi", func(c *fiber.Ctx) error { return c.SendString("hi") })
	h.Server.App.Get("/acc",h.Acc)
	h.Server.App.Get("/ppl",h.People)
	h.Server.App.Post("/t2s",h.T2S)
	h.Server.App.Post("/t2u",h.T2U)
}

