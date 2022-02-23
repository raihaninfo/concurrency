package model

const (
	ReqAdd        = iota
	ReqAvg        = iota
	ReqRandom     = iota
	ReqSpellCheck = iota
	ReqSearch     = iota
)
const ReqDataSize = 1 * 1024

type (
	ClientReq struct {
		Id      uint
		ReqType int
		Data    [ReqDataSize]byte
		Size    int
	}
)
