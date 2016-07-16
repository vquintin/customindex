package cache

import "sync"

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type result struct {
	value interface{}
	err   error
}

func newGoPLCache(f fct) *memo {
	return &memo{f: f, cache: make(map[interface{}]*entry)}
}

type memo struct {
	f     fct
	mu    sync.Mutex // guards cache
	cache map[interface{}]*entry
}

// Func is the type of the function to memoize.
type fct func(key interface{}) (interface{}, error)

func (memo *memo) get(key interface{}) (value interface{}, err error) {
	if hashable(key) {
		return memo.getFromCache(key)
	}
	return memo.f(key)
}

func (memo *memo) getFromCache(key interface{}) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}
