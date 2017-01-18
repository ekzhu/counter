package util

import "math"

// Counter for computing frequencies, discrete probabilities, and
// entropy over a collection of data values.
type Counter struct {
	counter    map[interface{}]int
	totalCount int
}

// Create a new Counter object.
func NewCounter() *Counter {
	return &Counter{
		counter: make(map[interface{}]int),
	}
}

// Update add a new element to the counter.
func (c *Counter) Update(elem interface{}) {
	if count, seen := c.counter[elem]; seen {
		c.counter[elem] = count + 1
	} else {
		c.counter[elem] = 1
	}
	c.totalCount++
}

// Freqs returns a slice of elements and a slice
// of corresponding integer frequencies.
func (c *Counter) Freqs() ([]interface{}, []int) {
	elems := make([]interface{}, 0, len(c.counter))
	counts := make([]int, 0, len(c.counter))
	for elem, count := range c.counter {
		elems = append(elems, elem)
		counts = append(counts, count)
	}
	return elems, counts
}

// Probs returns a slice of elements and a slice
// of corresponding discrete probabilities.
func (c *Counter) Probs() ([]interface{}, []float64) {
	elems := make([]interface{}, 0, len(c.counter))
	probs := make([]float64, 0, len(c.counter))
	for elem, count := range c.counter {
		elems = append(elems, elem)
		probs = append(probs, float64(count)/float64(c.totalCount))
	}
	return elems, probs
}

// Total returns the total number of elements counted.
func (c *Counter) Total() int {
	return c.totalCount
}

// Unique returns the number of unique elements counted.
func (c *Counter) Unique() int {
	return len(c.counter)
}

// Apply calls the function fn on each of the unique element
// counted. When an error is encountered in fn, it is immediately
// returned.
func (c *Counter) Apply(fn func(interface{}) error) error {
	for elem := range c.counter {
		if err := fn(elem); err != nil {
			return err
		}
	}
	return nil
}

// Entropy computes the entropy of the collection counted.
func (c *Counter) Entropy() float64 {
	var e float64
	for _, count := range c.counter {
		p := float64(count) / float64(c.totalCount)
		e -= p * math.Log(p)
	}
	return e
}

// PairCounter is for computing the co-occurrance frequencies, probailities
// and entropy of
// a pair of collections (i.e. two columns in a table).
type PairCounter struct {
	counter     map[interface{}](map[interface{}]int)
	totalCount  int
	uniqueCount int
}

// Create a new PairCounter object.
func NewPairCounter() *PairCounter {
	return &PairCounter{
		counter: make(map[interface{}](map[interface{}]int)),
	}
}

// Update adds a new pair of elements to the counter, one from
// each collection. The order of elements in the arguments
// must be consistent.
func (c *PairCounter) Update(elem1, elem2 interface{}) {
	if elem2Counter, seen := c.counter[elem1]; seen {
		if count, seen2 := elem2Counter[elem2]; seen2 {
			elem2Counter[elem2] = count + 1
		} else {
			elem2Counter[elem2] = 1
			c.uniqueCount++
		}
	} else {
		elem2Counter := make(map[interface{}]int)
		elem2Counter[elem2] = 1
		c.counter[elem1] = elem2Counter
		c.uniqueCount++
	}
	c.totalCount++
}

// Total returns the total number of pairs counted.
func (c *PairCounter) Total() int {
	return c.totalCount
}

// Unique returns the unique number of pairs counted.
func (c *PairCounter) Unique() int {
	return c.uniqueCount
}

// JointEntropy computes the joint entropy of the two collections
// counted.
func (c *PairCounter) JointEntropy() float64 {
	var e float64
	for _, elem2Counter := range c.counter {
		for _, count := range elem2Counter {
			p := float64(count) / float64(c.totalCount)
			e -= p * math.Log(p)
		}
	}
	return e
}
