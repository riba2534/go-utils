package utils

import (
	"sync"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		defer w.Done()
		cb()
	}()
}

func (w *WaitGroupWrapper) WrapWithArgs(cb func(args interface{}), args interface{}) {
	w.Add(1)
	go func(args interface{}) {
		defer w.Done()
		cb(args)
	}(args)
}
