package domain

type RequestDomain struct {
	RequestType uint8 `json:"reqtype"`
	BusinessID  int64 `json:"businessid"`
}
