package main

import (
	"GoConcurrentProgramming/chanbase/loadgen/lib"
	"time"
	"bytes"
	"strings"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/pkg/errors"
)

type ParamSet struct {
	Caller lib.Caller
	TimeoutNS time.Duration
	LPS uint32
	DurationNS time.Duration
	ResultCh chan *lib.CallResult
}

// Check 会检查当前值的所有字段的有效性。
// 若存在无效字段则返回值非nil。
func (pset *ParamSet) Check() error {
	var errMsgs []string
	if pset.Caller == nil {
		errMsgs = append(errMsgs, "Invalid caller!")
	}
	if pset.TimeoutNS == 0{
		errMsgs = append(errMsgs, "Invalid timeoutNS!")
	}
	if pset.LPS == 0{
		errMsgs = append(errMsgs, "Invalid lps(load per second)!")
	}
	if pset.DurationNS == 0{
		errMsgs = append(errMsgs, "Invalid durationNS!")
	}
	if pset.ResultCh == nil {
		errMsgs = append(errMsgs, "Invalid result channel!")
	}
	var buf bytes.Buffer
	buf.WriteString("Checking the parameters...")
	if errMsgs != nil{
		errMsg:=strings.Join(errMsgs, " ")
		buf.WriteString(fmt.Sprintf("NOT passed! (%s)", errMsg))
		log.Info(buf.String())
		return errors.New(errMsg)
	}
	buf.WriteString(fmt.Sprintf("Passed. (timeoutNS=%s, lps=%d, durationNS=%s)", pset.TimeoutNS, pset.LPS, pset.DurationNS))
	log.Info(buf.String())
	return nil
}