package worker

import (
	"context"
	"fmt"
	"testing"
)

const (
	workersCount = 4
	jobsCount    = 50
)

func TestPool(t *testing.T) {
	p := NewPool(workersCount)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go p.GenerateJobs(testJobs())

	go p.Run(ctx)

	for {
		select {
		case r, ok := <-p.Results():
			if !ok {
				continue
			}

			val, ok := r.Value.(int)
			if !ok {
				t.Fatalf("wrong type; expected int got %T", r.Value)
			}

			expectedVal := r.Descriptor.ID * r.Descriptor.ID
			if val != expectedVal {
				t.Fatalf("wrong value; expected %v got %v", expectedVal, val)
			}

		case <-p.Done():
			return
		}
	}
}

func testExecFn(ctx context.Context, arg interface{}) (interface{}, error) {
	argVal, ok := arg.(int)
	if !ok {
		return nil, fmt.Errorf("wrong argument type: expected int got %T", arg)
	}

	return argVal * argVal, nil
}

func testJobs() []Job {
	jobs := make([]Job, 0, jobsCount)

	for i := 0; i < jobsCount; i++ {
		jobs = append(jobs, Job{
			Descriptor: JobDescriptor{
				ID:       i,
				Metadata: nil,
			},
			Exec: testExecFn,
			Arg:  i,
		})
	}

	return jobs
}
