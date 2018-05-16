package main

import (
	"GoConcurrentProgramming/chanbase/loadgen/lib"
	"context"
	"time"
	"github.com/gpmgo/gopm/modules/log"
)

// 日志记录器。
//var logger = log.DLogger()

type myGenerator struct {
	caller lib.Caller
	timeoutNS time.Duration
	lps uint32
	durationNS time.Duration
	concurrency uint32
	tickets lib.GoTickets
	ctx context.Context
	cancelFunc context.CancelFunc
	callCount int64
	status uint32
	resultCh chan *lib.CallResult
}

func NewMyGenerator(pset ParamSet) lib.Generator {
	log.Info("New a load generator...")
	if err:=pset.Check(); err!=nil{
		return nil, err
	}
}

func main() {
	
}
