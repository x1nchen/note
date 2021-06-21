package topic_finallizer

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFinalizer(t *testing.T) {
	s := assert.New(t)

	w := NewCache()
	var cnt int = 0
	stopped := make(chan struct{})
	w.onStopped = func() {
		cnt++
		close(stopped)
	}

	s.Equal(0, cnt)

	w = nil

	runtime.GC()

	select {
	case <-stopped:
	case <-time.After(10 * time.Second):
		t.Fail()
	}

	s.Equal(1, cnt)
}
