package environment

import (
	"fmt"
	"sync"
)

var (
	environment Environment
	once        sync.Once
)

func Current() Environment {
	return environment
}

func Set(env string) (err error) {
	once.Do(func() {
		e, err := fromString(env)
		if err != nil {
			return
		}

		environment = e
	})

	return err
}

func fromString(s string) (Environment, error) {
	switch s {
	case Development.value:
		return Development, nil
	case Staging.value:
		return Staging, nil
	case Production.value:
		return Production, nil
	}

	return Unknown, fmt.Errorf("unknown environment: %s", s)
}

type Environment struct {
	value string
}

func (e Environment) String() string {
	return e.value
}

var (
	Unknown     = Environment{}
	Development = Environment{"development"}
	Staging     = Environment{"staging"}
	Production  = Environment{"production"}
)
