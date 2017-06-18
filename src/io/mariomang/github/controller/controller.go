package controller

import (
	"bytes"
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

	if c := r.Header.Get("Content-Type"); c != "application/json" {
		fmt.Println("Content-Type error")
		response := domain.NewErrorResponse(consts.ErrContentType, "ContentType error")
		w.Write(response.JsonBytes())
		return
	}

	jsonBytes := getJsonDataString(r)

	request := &domain.RequestDomain{}
	json.Unmarshal(jsonBytes, &request)

	fmt.Println(request)

	response := service.GenrateIDService(request)
	w.Write([]byte(response))
}

func getJsonDataString(r *http.Request) []byte {
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(result).Bytes()
}
