package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	f     Func
	cache map[string]*entry
	mu    sync.Mutex
}

// comment
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// comment
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// comment for Get
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}

func main() {
	// var hgb = httpGetBody
	var n sync.WaitGroup
	m := New(httpGetBody)
	incomingUrls := []string{"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
	}

	for _, url := range incomingUrls {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s, %s , %d bytes\n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}
