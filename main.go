package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Person struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {
	app := fiber.New()

	people := []Person{
		{FirstName: "Anders", LastName: "Jackson"},
		{FirstName: "Joe", LastName: "Rosenberg"},
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"status": "success",
			"people": people,
		})
	})

	log.Fatal(app.Listen(":3000"))

}
