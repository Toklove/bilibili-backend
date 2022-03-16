package Model

import "github.com/gofiber/fiber/v2"

type ResultItem struct {
	Code uint8  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func ApiResult(code uint8, msg string, data any) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(&ResultItem{
			Code: code,
			Msg:  msg,
			Data: data,
		})
	}
}
