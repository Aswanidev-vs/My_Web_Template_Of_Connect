package main

import (
	"log"
    "github.com/gofiber/fiber/v2"
)

func main() {
	app:=fiber.New()

	app.Post("/signup",SignupHandler)
	app.Post("/login",LoginHandler)
	log.Fatal(app.Listen(":1234"))
}