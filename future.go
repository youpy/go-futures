package futures

import (
	"reflect"
)

type Future chan interface{}

type Futurized func(in ...interface{}) Future

func New(fn interface{}) Futurized {
	if reflect.TypeOf(fn).NumOut() != 1 {
		panic("function must return single value")
	}

	return func(in ...interface{}) Future {
		var args []reflect.Value
		future := make(Future)

		for _, v := range in {
			args = append(args, reflect.ValueOf(v))
		}

		go func() {
			future <- reflect.ValueOf(fn).Call(args)[0].Interface()
		}()

		return future
	}
}
