package worker

import "context"

type Job struct {
	Descriptor JobDescriptor
	Exec       ExecFn
	Arg        interface{}
}

type JobDescriptor struct {
	JobID    int
	Metadata map[string]interface{}
}

type ExecFn func(ctx context.Context, arg interface{}) (interface{}, error)

type Result struct {
	Value      interface{}
	Err        error
	Descriptor JobDescriptor
}

func (j Job) execute(ctx context.Context) Result {
	value, err := j.Exec(ctx, j.Arg)
	if err != nil {
		return Result{
			Err:        err,
			Descriptor: j.Descriptor,
		}
	}

	return Result{
		Value:      value,
		Descriptor: j.Descriptor,
	}
}
