package internal

import "sync"

var instance = SafeCounter{requestCounter: 0}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mutex sync.Mutex
	requestCounter int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc() {
	c.mutex.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.requestCounter++
	c.mutex.Unlock()
}

// Dec decrements the counter for the given key.
func (c *SafeCounter) Dec() {
	c.mutex.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.requestCounter--
	c.mutex.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value() int {
	c.mutex.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mutex.Unlock()
	return c.requestCounter
}

func GetCounterInstance() *SafeCounter {
	return &instance
}

