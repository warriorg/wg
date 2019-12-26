package alg

import (
	"fmt"
	"testing"
	"time"
)

func Test_Window(t *testing.T) {
	win, err := New(10*time.Second, 2*time.Second)
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		win.Add()
		fmt.Println(win.Total())
	}
}
