package Router

import (
	"bilibili/application/Controller"
	"bilibili/core"
)

var app = core.App

func Routes() {
	app.Get("/", Controller.Index())
}
