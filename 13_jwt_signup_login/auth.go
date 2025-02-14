package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
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

		//now checking the user is alredy present in db or not
		db.Where("username = ?", user.Username).First(&user)
		if user.ID != 0 {
			return c.Status(fiber.StatusConflict).JSON( //409
				fiber.Map{
					"massage": "user already present",
				},
			)
		}
		//converting the password to the hashpassword
		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"eror": err.Error(),
			})
		}
		user.Password = string(hashed)
		db.Create(user) // here we creating the new user in the sqlite database
		//now We have to genarate the token
		//moving this method in anthor file
		token, err := GenerateJwtToken(user)
		if err != nil {
			log.Info("error", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to generate JWT token",
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
		dbUser := new(User)
		authUser := &User{
			Username: c.FormValue("username"),
			Password: c.FormValue("password"),
		}

		if authUser.Username == "" || authUser.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "username and password are required",
			})
		}
		//checking the user in db
		db.Where("username = ?", authUser.Username).First(&dbUser)
		//now the dbUser is actul prasentd db user

		//now check is it correct or not
		if dbUser.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(
				fiber.Map{
					"error": "user not found",
				},
			)
		}
		//if user present commper the password

		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(authUser.Password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{
					"error": "invalid credetials",
				},
			)
		}
		log.Info("User is Present in Db")
		token, err := GenerateJwtToken(dbUser)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				fiber.Map{
					"error": "Failed to generate JWT token",
				},
			)
		}

		c.Cookie(&fiber.Cookie{
			Name:   "jwt",
			Value:  token,
			Secure: !c.IsFromLocal(),

			HTTPOnly: !c.IsFromLocal(),
			MaxAge:   3600 * 24 * 7,
		},
		)
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"token": token,
			},
		)
	})
}
