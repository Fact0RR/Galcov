package handlers

import (
	"encoding/json"
	"net/http"
	"server/jwt"

	"github.com/gofiber/fiber/v2"
)



func (h* Handler) Acc(c *fiber.Ctx) error {
	//выводит все сбережения на аккаунте.
	token := c.Get("Token")

	claims := jwt.GetClaims(token)

	id := int(claims["id"].(float64))
	
	accData,err := h.Server.Store.Get_accaount(id)
	if err != nil{
		return c.SendStatus(http.StatusBadGateway)
	}
	jsonData, err:= json.Marshal(accData)
	if err != nil {
		return c.SendStatus(http.StatusBadGateway)
	}

	c.SendStatus(http.StatusOK)
	c.Set("Content-Type", "application/json")
	return c.Send(jsonData)
}