package worker

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
)

func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan Job, results chan<- Result) {
	defer wg.Done()

	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}

			results <- job.execute(ctx)
		case <-ctx.Done():
			logrus.Warnf("canceled worker, error details: %v\n", ctx.Err())
			results <- Result{
				Err: ctx.Err(),
			}
			return
		}
	}
}

type Pool struct {
	workersCount int
	jobs         chan Job
	results      chan Result
	done         chan struct{}
}

func NewPool(wCount int) Pool {
	return Pool{
		workersCount: wCount,
		jobs:         make(chan Job, wCount),
		results:      make(chan Result, wCount),
		done:         make(chan struct{}),
	}
}

func (p Pool) Run(ctx context.Context) {
	defer close(p.results)
	defer close(p.done)
	var wg sync.WaitGroup

	for i := 0; i < p.workersCount; i++ {
		wg.Add(1)
		go worker(ctx, &wg, p.jobs, p.results)
	}

	wg.Wait()
}

func (p Pool) Results() <-chan Result {
	return p.results
}

func (p Pool) Done() <-chan struct{} {
	return p.done
}

func (p Pool) GenerateJobs(jobs []Job) {
	defer close(p.jobs)

	for _, job := range jobs {
		p.jobs <- job
	}
}
