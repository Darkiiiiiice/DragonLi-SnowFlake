package main

import (
	"fmt"
	"github/mariomang/catrouter"
	"io/mariomang/github/controller"
	"time"
)

func main() {
	fmt.Println(time.Date(2017, 1, 1, 0, 0, 0, 0, time.Local).Unix())
	app := catrouter.NewDefaultApp("DragonLi", "MarioMang", "V-0.0.1", "")
	app.RegistController(catrouter.POST, "/primary/apply", controller.GenrateIDController)
	app.Run("", 9090)
}
