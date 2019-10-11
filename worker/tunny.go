package worker

import (
	"runtime"

	"github.com/Jeffail/tunny"
)

// CreateWorkerPool for http
func CreateWorkerPool(action func(interface{}) interface{}) *tunny.Pool {
	pool := tunny.NewFunc(runtime.GOMAXPROCS(runtime.NumCPU()), action)
	return pool
}
