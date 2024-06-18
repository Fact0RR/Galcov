package middleware

import (
	"net/http"
	"server/jwt"

	"github.com/gofiber/fiber/v2"
)

func LoginMiddleware(c *fiber.Ctx) error {	

	token := c.Get("Token")

	if jwt.CheckTokenOnValid(token){
		// Передача управления следующему обработчику
		err := c.Next()
		if err != nil{
			return c.SendStatus(http.StatusInternalServerError)
		}
	}else{
		return c.SendStatus(http.StatusUnauthorized)
	}

	return nil
}