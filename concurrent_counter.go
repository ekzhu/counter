package counter

import (
	cmap "github.com/orcaman/concurrent-map"
)

// ConcurrentCounter is a thread-safe version of Counter
type ConcurrentCounter struct {
	counter cmap.ConcurrentMap
}

func NewConcurrentCounter() *ConcurrentCounter {
	return &ConcurrentCounter{cmap.New()}
}

// Update add a new element to the counter.
func (c *ConcurrentCounter) Update(elem string) {
	c.counter.Upsert(elem, 1, func(exist bool, valueInMap interface{}, newValue interface{}) interface{} {
		if exist {
			return valueInMap.(int) + 1
		}
		return newValue
	})
}

// Has checks whether the elem has been counted before.
func (c *ConcurrentCounter) Has(elem string) bool {
	return c.counter.Has(elem)
}

// Unique returns the number of unique elements counted.
func (c *ConcurrentCounter) Unique() int {
	return c.counter.Count()
}

// Has checks whether the elem has been counted before.
func (c *ConcurrentCounter) Total() int {
	var total int
	for t := range c.counter.IterBuffered() {
		total += t.Val.(int)
	}
	return total
}
