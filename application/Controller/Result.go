package Controller

import (
	"bilibili/application/Model"
	"github.com/gofiber/fiber/v2"
)

func ApiResult(code uint8, msg string, data any) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(&Model.ResultItem{
			Code: code,
			Msg:  msg,
			Data: data,
		})
	}
}
