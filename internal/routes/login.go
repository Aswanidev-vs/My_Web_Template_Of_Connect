package routes

import "github.com/gofiber/fiber/v2"

type Login struct {
	Email string `json:email`
	Pwd   string `json:password`
}

var input Login

func LoginHandler(con *fiber.Ctx) error {
	if err := con.BodyParser(&input); err != nil {
		return con.Status(fiber.StatusBadRequest).SendString("invalid form data")
	}
}
