package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {

		fmt.Println("Someone hit me ")
		return c.SendString("The true measure of a shinobi is not how he lives, but how he dies. - Jiraiya👋!")
	})

	log.Fatal(app.Listen(":5555"))

}
