package domain

import "io/mariomang/github/consts"
import "encoding/json"

type ResponseDomain struct {
	ErrorCode uint8  `json:"errorcode"`
	ErrorMsg  string `json:"errormsg"`
	Data      string `json:"data"`
}

func NewResponseDomain() ResponseDomain {
	return ResponseDomain{
		ErrorCode: consts.Success,
		ErrorMsg:  "Success",
	}
}

func NewSuccessResponse(msg string, data string) ResponseDomain {
	return ResponseDomain{
		ErrorCode: consts.Success,
		ErrorMsg:  msg,
		Data:      data,
	}
}

func NewFailedResponse(msg string, data string) ResponseDomain {
	return ResponseDomain{
		ErrorCode: consts.Failure,
		ErrorMsg:  msg,
		Data:      data,
	}
}

func NewErrorResponse(code uint8, msg string, data string) ResponseDomain {
	return ResponseDomain{
		ErrorCode: code,
		ErrorMsg:  msg,
		Data:      data,
	}
}

func (r *ResponseDomain) JsonBytes() []byte {
	str, err := json.Marshal(r)
	if err != nil {
		return []byte(err.Error())
	}
	return str
}

func (r *ResponseDomain) JsonString() string {
	return string(r.JsonBytes())
}
