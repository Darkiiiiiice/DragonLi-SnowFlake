package main

import (
	"github/mariomang/catrouter"
	"io/mariomang/github/controller"
)

func main() {

	app := catrouter.NewDefaultApp("DragonLi", "MarioMang", "V-0.0.1", "")
	app.RegistController(catrouter.POST, "/primary/apply", controller.GenrateIDController)
	app.Run("", 9090)
}
