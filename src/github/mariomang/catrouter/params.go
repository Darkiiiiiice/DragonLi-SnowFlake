package catrouter

import (
	"github.com/julienschmidt/httprouter"
)

type Params struct {
	params httprouter.Params
}

func (ps Params) ByName(name string) string {
	return ps.params.ByName(name)
}
