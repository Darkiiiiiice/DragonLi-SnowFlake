package service

import (
	"encoding/json"
	"io/mariomang/github/domain"
	"io/mariomang/github/snowflake"
)

func GenrateIDService(request *domain.RequestDomain) string {
	sf := snowflake.NewSnowFlake(request.WorkID, request.MachineID)
	id := sf.GetID()
	response, _ := json.Marshal(domain.NewSuccessResponse("Success", request.WorkID, id))
	return string(response)
}
