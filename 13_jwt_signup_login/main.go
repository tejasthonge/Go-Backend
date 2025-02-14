package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func main() {

	//initialze the database

	db := InitDb()
	fmt.Println(db)
	//creating the fiber app Instance
	app := fiber.New(fiber.Config{AppName: "Auth API"})

	// public routes
	// app.Get("/", func(c fiber.Ctx) error {
	// 	return c.SendString("Starting the server..")
	// })
	// authRouter := app.Group("/auth")

	AuthHandlers(app.Group("/auth"), db)

	//start the server on port 8055
	app.Listen(":8055")

}
