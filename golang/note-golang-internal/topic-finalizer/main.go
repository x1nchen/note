package topic_finallizer

import (
	"runtime"
	"time"
)

type Cache = *wrapper

type wrapper struct {
	*cache
}

type cache struct {
	content   string
	stopCH    chan struct{}
	onStopped func()
}

func newCache() *cache {
	return &cache{
		content: "some thing",
		stopCH:  make(chan struct{}),
	}
}

func NewCache() Cache {
	w := &wrapper{
		cache: newCache(),
	}
	go w.cache.run()
	runtime.SetFinalizer(w, (*wrapper).stop)
	runtime.KeepAlive(w)
	return w
}

func (w *wrapper) stop() {
	w.cache.stop()
}

func (c *cache) run() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// do some thing
		case <-c.stopCH:
			if c.onStopped != nil {
				c.onStopped()
			}
			return
		}
	}
}

func (c *cache) stop() {
	close(c.stopCH)
}


