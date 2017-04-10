package counter

import (
	"sync"
	"testing"
)

func TestConcurrentCounter(t *testing.T) {
	testData := []string{"a", "b", "b", "c", "c", "d", "e", "e", "e", "f"}
	counter := NewConcurrentCounter()

	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func() {
			for _, v := range testData {
				counter.Update(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	if counter.Total() != len(testData)*4 {
		t.Error("Incorrect total count")
	}

	c := counter.Unique()
	if c != 6 {
		t.Error(c)
	}
}
