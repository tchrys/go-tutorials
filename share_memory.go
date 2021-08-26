package main

import (
	"log"
	"net/http"
	"time"
)

const (
	numPollers = 2 // numbers of poll goroutines to launch
	pollInterval = 60 * time.Second // how often to poll each URL
	statusInterval = 10 * time.Second // how often to log status
	errTimeout = 10 * time.Second // back-off timeout on error
)

var urls = [] string {
	"http://www.google.com/",
	"http://golang.org/",
	"http://blog.golang.org",
}

// the last-known state of a URL
type State struct {
	url string
	status string
}

func StateMonitor(updateInverval time.Duration) chan<- State {
	updates := make(chan State)
	urlStatus := make(map[string]string)
	ticker := time.NewTicker(updateInverval)
	go func() {
		for {
			select {
			case <-ticker.C:
				logState(urlStatus)
			case s := <- updates:
				urlStatus[s.url] = s.status
			}
		}
	}()
	return updates
}

func logState(s map[string]string) {
	log.Println("Current state: ")
	for k, v := range s {
		log.Printf(" %s %s", k, v)
	}
}

type Resource struct {
	url string
	errCount int
}

func (r * Resource) Poll() string {
	resp, err := http.Head(r.url)
	if err != nil {
		log.Println("Error", r.url, err)
		r.errCount++
		return err.Error()
	}
	r.errCount = 0
	return resp.Status
}

func (r * Resource) Sleep(done chan<- *Resource) {
	time.Sleep(pollInterval + errTimeout * time.Duration(r.errCount))
	done <- r
}

func Poller(in <-chan *Resource, out chan<- *Resource, status chan<- State) {
	for r := range in {
		s := r.Poll()
		status <- State {r.url, s}
		out <- r
	}
}

func main() {
	pending, complete := make(chan *Resource), make(chan *Resource)
	status := StateMonitor(statusInterval)
	for i := 0; i < numPollers; i++ {
		go Poller(pending, complete, status)
	}
	go func() {
		for _, url := range urls {
			pending <- &Resource{url: url}
		}
	}()
	for r := range complete {
		go r.Sleep(pending)
	}
}


