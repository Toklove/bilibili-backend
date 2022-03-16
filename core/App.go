package core

import (
	"bilibili/application/Config"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

var App = fiber.New()
var Conf = Config.Config{}

func RunApi() {
	defer func() {
		if err := App.Listen(fmt.Sprintf(":%v", Conf.GetConf().Server.Port)); err != nil {
			panic("端口被占用请重试")
		}
	}()
}
