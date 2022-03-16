package Controller

import (
	"bilibili/application/Model"
	"github.com/gofiber/fiber/v2"
)

func Index() func(c *fiber.Ctx) error {
	return Model.ApiResult(200, "success", nil)
}
