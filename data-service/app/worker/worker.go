package worker

import (
	"runtime"
	"sync"
)

type Pool struct {
	pool *sync.Pool
}

func NewPool() *Pool {
	numCores := runtime.NumCPU()
	workerPool := &sync.Pool{
		New: func() any {

		},
	}
}

type Worker struct {
	
}
