package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type TransferToSelf struct{
	Id_user int `json:"id_user"`
	Id_val_old int `json:"id_val_old"`
	Id_val_new int `json:"id_val_new"`
	Count float64 `json:"count"`
}

func (h *Handler) T2S(c *fiber.Ctx) error {
	print("T2S")
	var user TransferToSelf

	// Привязываем тело запроса к структуре Message
	if err := c.BodyParser(&user); err != nil {
		print(err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}
	fmt.Println(user)
	if err := h.Server.Store.TransferToSelf(user.Id_user,user.Id_val_old,user.Id_val_new,user.Count);err!= nil{
		print(err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}

	
	return c.SendStatus(http.StatusOK)
}




type TransferToUser struct{
	Id_old int `json:"id_old"`
	Id_new int `json:"id_new"`
	Id_val int `json:"id_val"`
	Count float64 `json:"count"`
}

func (h *Handler) T2U(c *fiber.Ctx) error{
	var user TransferToUser

	// Привязываем тело запроса к структуре Message
	if err := c.BodyParser(&user); err != nil {
		print(err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}
	
	if err := h.Server.Store.TransferToUser(user.Id_old,user.Id_new,user.Id_val,user.Count);err!= nil{
		print(err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}

	
	return c.SendStatus(http.StatusOK)
}