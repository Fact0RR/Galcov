package handlers

import (
	"net/http"
	"server/jwt"

	"github.com/gofiber/fiber/v2"
)



func (h* Handler)  Auth(c *fiber.Ctx) error {
	var user User
	println("Auth")
	// Привязываем тело запроса к структуре Message
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	// Проверка пароля в бд
	is_passed, err := h.Server.Store.Check_user(user.Login,user.Password)
	if err != nil{
		print(err.Error())
		return c.SendStatus(http.StatusInternalServerError)
	}
	if  is_passed == 0 {
		return c.SendStatus(http.StatusUnauthorized)
	}

	token,err := jwt.GenerateJWT(is_passed,user.Login)
	if err != nil{
		return c.SendStatus(http.StatusInternalServerError)
	}
	
	return c.SendString(token)

}