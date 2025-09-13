package goc

import (
	"runtime"
)

func Go[T any](target *T, f func(chan struct{})) {
	exit := make(chan struct{})

	go func() {
		f(exit)
	}()
	cleanUp := func(exit chan struct{}) {
		close(exit)
	}
	runtime.AddCleanup(target, cleanUp, exit)
}
