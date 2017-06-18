package domain

import "io/mariomang/github/consts"
import "encoding/json"

type ResponseDomain struct {
	ErrorCode uint8  `json:"errorcode"`
	ErrorMsg  string `json:"errormsg"`
	WorkID    int64  `json:"workid"`
	GlobalID  int64  `json:"globalid"`
}

func NewResponseDomain() ResponseDomain {
	return ResponseDomain{
		ErrorCode: consts.Success,
		ErrorMsg:  "Success",
	}
}

func NewSuccessResponse(msg string, workid int64, globalid int64) ResponseDomain {
	return ResponseDomain{
		ErrorCode: consts.Success,
		ErrorMsg:  msg,
		WorkID:    workid,
		GlobalID:  globalid,
	}
}

func NewFailedResponse(msg string) ResponseDomain {
	return ResponseDomain{
		ErrorCode: consts.Failure,
		ErrorMsg:  msg,
	}
}

func NewErrorResponse(code uint8, msg string) ResponseDomain {
	return ResponseDomain{
		ErrorCode: code,
		ErrorMsg:  msg,
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
