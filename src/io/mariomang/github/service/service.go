package service

import (
	"encoding/json"
	"fmt"
	"io/mariomang/github/consts"
	"io/mariomang/github/domain"
	"io/mariomang/github/snowflake"
)

func GenrateIDService(request *domain.RequestDomain) string {
	sf := snowflake.NewSnowFlake(request.WorkID, request.MachineID)
	id := sf.GetID()
	response, err := json.Marshal(domain.NewSuccessResponse("Success", request.WorkID, id))
	if err != nil {
		emsg := fmt.Sprintf(consts.JsonMarshalErrorMsg, err)
		fmt.Println(emsg)
		return emsg
	}
	return string(response)
}
