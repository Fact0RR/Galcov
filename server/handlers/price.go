package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Price(c *fiber.Ctx) error {

	arrPice, err := h.Server.Store.Get_price()
	if err != nil{
		return c.SendStatus(http.StatusBadGateway)
	}
	jsonData, err:= json.Marshal(arrPice)
	if err != nil {
		return c.SendStatus(http.StatusBadGateway)
	}
	c.SendStatus(http.StatusOK)
	c.Set("Content-Type", "application/json")
	return c.Send(jsonData)
}