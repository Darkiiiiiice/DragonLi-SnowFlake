package consts

import (
	"flag"
	"os"
)

//项目相关常量
//AppName 项目名称
//Author  作者
//Version 版本号
//Email   邮箱
const (
	AppName = "DragonLi"
	Author  = "MarioMang"
	Version = "v-0.0.1"
	Email   = "853275286@qq.com"
)

// ResponseDomain ErrorCode
// Success success response errorcode
// Failure failure response errorcode
// ContentTypeErrorCOde cannot match the content-type response errorcode
// PostDataErrorCode receive post request data response errorcode
// JSONUnmarshalErrorCode use standard libary json.Unmarshal(v interface{}) response errorcode
const (
	Success = iota
	Failure
	ContentTypeErrorCode
	PostDataErrorCode
	JSONUnmarshalErrorCode
)

// ResponseDomain ErrorMsg
// Success success response errormsg
// Failure failure response errormsg
// ContentTypeErrorMsg cannot match the content-type response errormsg
// PostDataErrorMsg receive psot request data response errormsg
// JSONUnmarshalErrorMSg use standard libary json.Unmarshal(v interface{}) response errormsg
// JSONMarshalErrorMsg use standard libary json.Marshal() respone errormsg
const (
	SuccessMsg            = "Success"
	FailureMsg            = "Failure"
	ContentTypeErrorMsg   = "content-type match error: %s\n"
	PostDataErrorMsg      = "cannot get post data: %v\n"
	JSONUnmarshalErrorMsg = "json unmarshal() error: %v\n"
	JSONMarshalErrorMsg   = "json marshal() error: %v\n"
)

// SnowFlake constants
// BaseTime base timestamp default: 2017-1-1 0:0:0
// WorkIDBits workId of bits
// MachineIDBits machine of bits
// SequenceBits sequence of bits
const (
	BaseTime      int64 = 1483200000
	WorkIDBits    int64 = 6
	MachineIDBits int64 = 6
	SequenceBits  int64 = 10
)

// ContentType need to verify the ContentType
const (
	ContentType = "application/json"
)

// Program flag parameter
// ListenIP the server will use this ip to listen
// ListenPort the server will use this port to listen
// Help use this parameter print the help message
var (
	ListenIP   string
	ListenPort int
	Help       bool
)

func init() {
	flag.StringVar(&ListenIP, "h", "127.0.0.1", " listening host")
	flag.IntVar(&ListenPort, "p", 8085, " listening port")
	flag.BoolVar(&Help, "-help", false, " help")
	flag.Parse()

	if Help {
		flag.Usage()
		os.Exit(1)
	}

	if !flag.Parsed() {
		flag.Usage()
		os.Exit(1)
	}
}
