package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡Bienvenido a la calculadora!")
	})

	app.Get("/sum/:a/:b", func(c *fiber.Ctx) error {
		a, err1 := strconv.Atoi(c.Params("a"))
		b, err2 := strconv.Atoi(c.Params("b"))
		if err1 != nil || err2 != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid parameters")
		}
		result := a + b
		return c.SendString(fmt.Sprintf("Sum: %d", result))
	})

	app.Get("/subtract/:a/:b", func(c *fiber.Ctx) error {
		a, err1 := strconv.Atoi(c.Params("a"))
		b, err2 := strconv.Atoi(c.Params("b"))
		if err1 != nil || err2 != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid parameters")
		}
		result := a - b
		return c.SendString(fmt.Sprintf("Subtraction: %d", result))
	})

	app.Get("/multiply/:a/:b", func(c *fiber.Ctx) error {
		a, err1 := strconv.Atoi(c.Params("a"))
		b, err2 := strconv.Atoi(c.Params("b"))
		if err1 != nil || err2 != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid parameters")
		}
		result := a * b
		return c.SendString(fmt.Sprintf("Multiplication: %d", result))
	})

	app.Get("/divide/:a/:b", func(c *fiber.Ctx) error {
		a, err1 := strconv.Atoi(c.Params("a"))
		b, err2 := strconv.Atoi(c.Params("b"))
		if err1 != nil || err2 != nil || b == 0 {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid parameters")
		}
		result := float64(a) / float64(b)
		return c.SendString(fmt.Sprintf("Division: %.2f", result))
	})

	// Run server on port 3000
	app.Listen(":3000")
}
