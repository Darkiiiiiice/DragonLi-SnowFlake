package catrouter

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type DefaultRouter struct {
	router *httprouter.Router
}

type Controller func(w http.ResponseWriter, r *http.Request, p *Params)

func (dr *DefaultRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	path := r.URL.Path
	dr.router.ServeHTTP(w, r)
	endTime := time.Now()
	fmt.Printf("path:%s time:%s\n", path, endTime.Sub(startTime))
}

func NewDefaultRouter() *DefaultRouter {
	return &DefaultRouter{
		router: httprouter.New(),
	}
}

func (r *DefaultRouter) RegistRouter(method, path string, controller Controller) {
	r.router.Handle(method, path,
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			params := &Params{params: p}
			controller(w, r, params)
		},
	)
}

func (r *DefaultRouter) GET(path string, controller Controller) {
	r.router.GET(path,
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			params := &Params{params: p}
			controller(w, r, params)
		},
	)
}

func (r *DefaultRouter) POST(path string, controller Controller) {
	r.router.POST(path,
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			params := &Params{params: p}
			controller(w, r, params)
		},
	)
}

func (r *DefaultRouter) DELETE(path string, controller Controller) {
	r.router.DELETE(path,
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			params := &Params{params: p}
			controller(w, r, params)
		},
	)
}

func (r *DefaultRouter) HEAD(path string, controller Controller) {
	r.router.HEAD(path,
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			params := &Params{params: p}
			controller(w, r, params)
		},
	)
}

func (r *DefaultRouter) PUT(path string, controller Controller) {
	r.router.PUT(path,
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			params := &Params{params: p}
			controller(w, r, params)
		},
	)
}

func (r *DefaultRouter) OPTIONS(path string, controller Controller) {
	r.router.OPTIONS(path,
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			params := &Params{params: p}
			controller(w, r, params)
		},
	)
}

func (r *DefaultRouter) PATCH(path string, controller Controller) {
	r.router.PATCH(path,
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			params := &Params{params: p}
			controller(w, r, params)
		},
	)
}

func (r *DefaultRouter) ServeFiles(path string, root http.FileSystem) {
	r.router.ServeFiles(path, root)
}

func (r *DefaultRouter) HandlerFunc(method, path string, handler http.HandlerFunc) {
	r.router.Handler(method, path, handler)
}

func (r *DefaultRouter) Handler(method, path string, handler http.Handler) {
	r.RegistRouter(method, path,
		func(w http.ResponseWriter, r *http.Request, _ *Params) {
			handler.ServeHTTP(w, r)
		},
	)
}
