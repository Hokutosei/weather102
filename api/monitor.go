package main

import (
	"fmt"
	"runtime"
	"time"

	utils "github.com/Hokutosei/hokutoseiUtils"
)

func reportGoRoutines(reportDelaySec int) {
	for t := range time.Tick(time.Duration(reportDelaySec) * time.Second) {
		_ = t
		utils.Info(fmt.Sprintf("currently have goroutines -->> %v", runtime.NumGoroutine()))
	}
}
