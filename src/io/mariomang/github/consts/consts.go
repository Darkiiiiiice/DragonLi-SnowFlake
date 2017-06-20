package consts

const (
	AppName = "DragonLi"
	Author  = "MarioMang"
	Version = "v-0.0.1"
	Email   = "853275286@qq.com"
)

const (
	ListenIP   = "127.0.0.1"
	ListenPort = 9090
)

const (
	Success = iota
	Failure
	ContentTypeErrorCode
	PostDataErrorCode
	JsonUnmarshalErrorCode
)

const (
	BaseTime      int64 = 1483200000 //2017-1-1 0:0:0
	WorkIDBits    int64 = 6
	MachineIDBits int64 = 6
	SequenceBits  int64 = 10
)

const (
	ContentType = "application/json"
)

const (
	ContentTypeErrorMsg   = "content-type match error: %s\n"
	PostDataErrorMsg      = "cannot get post data: %v\n"
	JsonUnmarshalErrorMsg = "json unmarshal() error: %v\n"
	JsonMarshalErrorMsg   = "json marshal() error: %v\n"
)
