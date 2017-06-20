package controller

import (
	"encoding/json"
	"fmt"
	"github/mariomang/catrouter"
	"io/ioutil"
	"io/mariomang/github/consts"
	"io/mariomang/github/domain"
	"io/mariomang/github/service"
	"net/http"
)

func GenrateIDController(w http.ResponseWriter, r *http.Request, p *catrouter.Params) {

	if c := r.Header.Get("Content-Type"); c != consts.ContentType {
		emsg := fmt.Sprintf(consts.ContentTypeErrorMsg, c)
		fmt.Println(emsg)
		response := domain.NewErrorResponse(consts.ContentTypeErrorCode, emsg)
		w.Write(response.JsonBytes())
		return
	}

	jsonBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		emsg := fmt.Sprintf(consts.PostDataErrorMsg, err)
		fmt.Println(emsg)
		response := domain.NewErrorResponse(consts.PostDataErrorCode, emsg)
		w.Write(response.JsonBytes())
		return
	}

	request := new(domain.RequestDomain)
	err = json.Unmarshal(jsonBytes, &request)
	if err != nil {
		emsg := fmt.Sprintf(consts.JsonUnmarshalErrorMsg, err)
		fmt.Println(emsg)
		response := domain.NewErrorResponse(consts.JsonUnmarshalErrorCode, emsg)
		w.Write(response.JsonBytes())
		return
	}

	response := service.GenrateIDService(request)
	w.Write([]byte(response))
}
