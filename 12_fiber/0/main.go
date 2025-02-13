package main

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type User struct {
	Status string `json:"status"`
	Id     int64  `json:"id"`
	Name   string `json:"name"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Jay Shree Ram")

	})
	app.Get("/api/users/get/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		intId, err := strconv.ParseInt(id, 10, 64)
		 
		if err != nil {
			c.Status(fiber.ErrBadRequest.Code)
			return c.SendString("Plese pass Valid id")
		}
		user := User{
			Status: "OK",
			Id:     intId,
			Name:   "Tejas",
		}
		return c.JSON(&user)
	})

	app.Listen(":8055")
}
