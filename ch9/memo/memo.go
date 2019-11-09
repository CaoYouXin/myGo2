package main

import (
	"fmt"
	"time"
)

type fn func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type request struct {
	key      string
	response chan<- result
	cancel   chan struct{}
}

type memo struct {
	requests chan request
}

func (m *memo) get(key string, cancel chan struct{}) (interface{}, error) {
	response := make(chan result)
	m.requests <- request{key, response, cancel}
	res := <-response
	return res.value, res.err
}

func (m *memo) close() { close(m.requests) }

func (m *memo) server(f fn) {
	cache := make(map[string]*entry)
	for req := range m.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}

		select {
		case <-req.cancel:
			delete(cache, req.key)
		default:
			go e.deliver(req.response)
		}
	}
}

func (e *entry) call(f fn, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}

func newMemo(f fn) *memo {
	m := &memo{requests: make(chan request)}
	go m.server(f)
	return m
}

func t(key string) (interface{}, error) {
	time.Sleep(time.Second)
	return fmt.Sprint(key, time.Now()), nil
}

func main() {
	m := newMemo(t)
	cancel := make(chan struct{})
	go func() {
		fmt.Println(m.get("Hello", cancel))
	}()
	go func() {
		time.Sleep(500 * time.Millisecond)
		cancel <- struct{}{}
	}()
}
