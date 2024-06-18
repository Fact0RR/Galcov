package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h* Handler)  Reg(c *fiber.Ctx) error {
	var user User

	// Привязываем тело запроса к структуре Message
	if err := c.BodyParser(&user); err != nil {
		print(err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}
	
	if err := h.Server.Store.Add_user(user.Login,user.Password);err!= nil{
		print(err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}

	c.SendStatus(http.StatusOK)
	return c.SendString("reg_succesful")

}