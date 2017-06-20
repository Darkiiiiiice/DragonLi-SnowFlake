package main

import (
	"github/mariomang/catrouter"
	"io/mariomang/github/consts"
	"io/mariomang/github/controller"
)

//The SnowFlake program enter
func main() {
	app := catrouter.NewDefaultApp(consts.AppName, consts.Author, consts.Version, consts.Email)
	app.RegistController(catrouter.POST, "/primary/apply", controller.GenrateIDController)
	app.Run(consts.ListenIP, consts.ListenPort)
}
