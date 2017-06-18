package domain

type RequestDomain struct {
	RequestType uint8 `json:"reqtype"`
	WorkID      int64 `json:"workid"`
}
