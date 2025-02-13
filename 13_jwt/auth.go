package main

import (
	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AuthHandlers(route fiber.Router, db *gorm.DB) {
	route.Post("/register", func(c fiber.Ctx) error {

		user := &User{
			Username: c.FormValue("username"),
			Password: c.FormValue("password"),
		}
		if user.Username == "" || user.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(
				fiber.Map{"error": "user name and passwad are must required"},
			)
		}
		//converting the password to the hashpassword
		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"eror": err.Error(),
			})
		}
		user.Password = string(hashed)
		db.Create(user) // here we creating the new user in the sqlite database
		//now We have to genarate the token
		//moving this method in anthor file
		token, err := GenarateJwtToken(user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			HTTPOnly: !c.IsFromLocal(),
			Secure:   !c.IsFromLocal(),
			MaxAge:   3600 * 24 * 7, //7 days
		})
		return c.Status(fiber.StatusCreated).JSON(
			fiber.Map{
				"token": token,
			},
		)
	})

	route.Post("/login", func(c fiber.Ctx) error {
		return nil
	})
}
