package lib

import "time"

type CallResult struct {
	ID int64
	Req RawReq
	Resp RawResp
	Code RetCode
	Msg string
	Elapse time.Duration
}

type RawReq struct {
	ID int64
	Req []byte
}

type RawResp struct {
	ID int64
	Resp []byte
	Err error
	Elapse time.Duration
}

// RetCode 表示结果代码的类型。
type RetCode int

// 保留 1 ~ 1000 给载荷承受方使用。
const (
	RET_CODE_SUCCESS              RetCode = 0    // 成功。
	RET_CODE_WARNING_CALL_TIMEOUT         = 1001 // 调用超时警告。
	RET_CODE_ERROR_CALL                   = 2001 // 调用错误。
	RET_CODE_ERROR_RESPONSE               = 2002 // 响应内容错误。
	RET_CODE_ERROR_CALEE                  = 2003 // 被调用方（被测软件）的内部错误。
	RET_CODE_FATAL_CALL                   = 3001 // 调用过程中发生了致命错误！
)

type Generator interface {
	Start() bool
	Stop() bool
	Status() uint32
	CallCount() int64
}