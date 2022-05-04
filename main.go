package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Person struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func main() {
	app := fiber.New()
	fmt.println()

	people := []Person{
		{FirstName: "Anders", LastName: "Jackson", Email: "jackson@mail.com"},
		{FirstName: "Joe", LastName: "Rosenberg", Email: "joe@email.com"},
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"status": "success",
			"people": people,
		})
	})

	log.Fatal(app.Listen(":3000"))

}
