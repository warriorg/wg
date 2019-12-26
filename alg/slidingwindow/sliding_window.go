package alg

import (
	"errors"
	"sync"
	"time"
)

// Window window
type Window struct {
	sync.RWMutex

	window   time.Duration
	tick     time.Duration
	samples  []int64
	stopOnce sync.Once
	stopping chan struct{}
	pos      int
	size     int
}

// New new window
// New(time.Second*5, time.Second)
func New(window, tick time.Duration) (*Window, error) {
	if window == 0 {
		return nil, errors.New("sliding window cannot be zero")
	}

	if tick == 0 {
		return nil, errors.New("tick cannot be zero")
	}

	if window <= tick || window%tick != 0 {
		return nil, errors.New("window size has to be a multiplier of granularity size")
	}

	win := &Window{
		window:   window,
		tick:     tick,
		samples:  make([]int64, int(window/tick)),
		stopping: make(chan struct{}, 1),
	}

	go win.shifter()

	return win, nil
}

func (win *Window) shifter() {
	ticker := time.NewTicker(win.tick)

	for {
		select {
		case <-ticker.C:
			win.slidingWindow()
		case <-win.stopping:
			return
		}
	}
}

func (win *Window) slidingWindow() {
	win.Lock()
	defer win.Unlock()

	win.pos = win.pos + 1
	if win.pos >= len(win.samples) {
		win.pos = 0
	}
	win.samples[win.pos] = 0
}

// Add increments the value of the current sample.
func (win *Window) Add() {
	win.Lock()
	defer win.Unlock()
	win.samples[win.pos]++
}

// AddCount increments the value of the current sample.
func (win *Window) AddCount(n int64) {
	win.Lock()
	defer win.Unlock()
	win.samples[win.pos] += n
}

// Reset the samples in this sliding time window.
func (win *Window) Reset() {
	win.Lock()
	defer win.Unlock()

	win.pos, win.size = 0, 0
	for i := range win.samples {
		win.samples[i] = 0
	}
}

// Total returns the sum of all values over the window
func (win *Window) Total() int64 {

	win.RLock()
	defer win.RUnlock()

	var total int64
	for i := range win.samples {
		total += win.samples[i]
	}
	return total
}

// Stop the shifter of this sliding time window. A stopped SlidingWindow cannot
// be started again.
func (win *Window) Stop() {
	win.stopOnce.Do(func() {
		win.stopping <- struct{}{}
	})
}
