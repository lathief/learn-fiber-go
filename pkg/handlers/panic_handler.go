package handlers

import "github.com/gofiber/fiber/v2"

func PanicHandler(c *fiber.Ctx, e interface{}) {
	err := c.Status(500).JSON("Internal Server Error")
	if err != nil {
		return
	}
}
