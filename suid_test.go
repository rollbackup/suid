package suid

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSuid(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		for i := 0; i < 1025; i++ {
			fmt.Println(Generate(i))
		}
	}()
	wg.Wait()
}
