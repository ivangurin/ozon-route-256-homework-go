package closer

import (
	"os"
	"os/signal"
	"sync"

	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

type ICloser interface {
	Add(f ...func() error)
	Wait()
	Signal()
	CloseAll()
}

type closer struct {
	sync.Mutex
	once     sync.Once
	done     chan struct{}
	funcs    []func() error
	shutdown chan os.Signal
}

// os.Interrupt, syscall.SIGINT, syscall.SIGTERM
func NewCloser(sig ...os.Signal) ICloser {
	c := &closer{
		done:     make(chan struct{}),
		shutdown: make(chan os.Signal, 1),
	}

	if len(sig) > 0 {
		go func() {
			signal.Notify(c.shutdown, sig...)
			<-c.shutdown
			signal.Stop(c.shutdown)
			c.CloseAll()
		}()
	}

	return c
}

func (c *closer) Add(f ...func() error) {
	c.Lock()
	c.funcs = append(c.funcs, f...)
	c.Unlock()
}

func (c *closer) Wait() {
	<-c.done
}

func (c *closer) Signal() {
	close(c.shutdown)
}

func (c *closer) CloseAll() {
	c.once.Do(func() {
		logger.Info("Graceful shutdown started...")
		defer logger.Info("Graceful shutdown finished")

		defer close(c.done)

		c.Lock()
		funcs := c.funcs
		c.Unlock()

		for i := len(funcs) - 1; i >= 0; i-- {
			err := c.funcs[i]()
			if err != nil {
				logger.Errorf("failed to close some func from shutdown: %v", err)
			}
		}
	})
}
