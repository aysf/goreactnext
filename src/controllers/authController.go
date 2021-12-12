package controllers

import (
	"github.com/aysf/goreactnext/src/database"
	"github.com/aysf/goreactnext/src/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password do not match",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

	user := models.User{
		FirstName:   data["first_name"],
		LastName:    data["last_name"],
		Email:       data["email"],
		Password:    string(password),
		IsAmbasador: false,
	}

	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"message": "hello",
	})

}
