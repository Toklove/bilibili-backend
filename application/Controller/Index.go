package Controller

import (
	"github.com/gofiber/fiber/v2"
)

func Index() func(c *fiber.Ctx) error {
	return ApiResult(200, "success", nil)
}
