package errgroup

import (
	"context"
	"sync"
)

type errgroup struct {
	ctx    context.Context
	cancel func()
	done   chan error
	funcs  []func() error
	queue  chan struct{}
}

func NewErrGroup(ctx context.Context, workers int) (*errgroup, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &errgroup{
		ctx:    ctx,
		cancel: cancel,
		done:   make(chan error),
		queue:  make(chan struct{}, workers),
	}, ctx
}

func (eg *errgroup) Go(f func() error) {
	eg.funcs = append(eg.funcs, f)
}

func (eg *errgroup) Wait() error {
	go eg.run()
	return <-eg.done
}

func (eg *errgroup) run() {
	defer close(eg.done)
	var err error
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for _, f := range eg.funcs {
		eg.queue <- struct{}{}
		if err != nil {
			break
		}
		wg.Add(1)
		go func() {
			defer func() {
				<-eg.queue
				wg.Done()
			}()
			lerr := f()
			if lerr != nil {
				if err == nil {
					mu.Lock()
					if err == nil {
						err = lerr
						eg.cancel()
					}
					mu.Unlock()
				}
			}
		}()
	}
	wg.Wait()
	eg.done <- err
}
