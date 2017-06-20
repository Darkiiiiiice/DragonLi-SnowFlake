package catrouter

import (
	"fmt"
	"net/http"
)

// DefaultApp 默认实现的App结构
// DefaultApp.Name    App-名称
// DefaultApp.Author  App-作者
// DefaultApp.Version App-版本号
// DefaultApp.Email   App-Email
// DefaultApp.IP      App-监听IP
// Default.Port       App-监听端口
type DefaultApp struct {
	Name    string
	Author  string
	Version string
	Email   string
	IP      string
	Port    int
	router  *DefaultRouter
}

const (
	DELETE  = "DELETE"
	GET     = "GET"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	PATCH   = "PATCH"
	POST    = "POST"
	PUT     = "PUT"
)

// NewDefaultApp 创建一个新的默认App
// @Param name string Application name
// @Param author string Application author
func NewDefaultApp(name string, author string, version string, email string) *DefaultApp {
	return &DefaultApp{
		Name:    name,
		Author:  author,
		Version: version,
		Email:   email,
		router:  NewDefaultRouter(),
	}
}

func (app *DefaultApp) Run(ip string, port int) {

	app.IP = ip
	app.Port = port

	// defaultRouter := NewDefaultRouter()
	addr := fmt.Sprintf("%s:%d", app.IP, app.Port)

	fmt.Printf("======================== %s start ========================\n", app.Name)
	fmt.Printf("=== Author: %s \n", app.Author)
	fmt.Printf("=== Email:  %s \n", app.Email)
	fmt.Printf("=== Version: %s \n", app.Version)
	fmt.Printf("=== Listening: %s:%d \n", app.IP, app.Port)

	http.ListenAndServe(addr, app.router)
}

func (app *DefaultApp) RegistController(method string, path string, controller Controller) {
	app.router.RegistRouter(method, path, controller)
}
