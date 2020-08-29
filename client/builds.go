package main

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/r3labs/sse"
	"github.com/workhorse/api"
)

type WatchHandler func(obj interface{})

type Builds struct {
}

func (b *Builds) Watch(url string, handler WatchHandler) error {

	client := sse.NewClient(url)

	err := client.Subscribe("build", func(msg *sse.Event) {
		build := &api.Build{}
		json.Unmarshal(msg.Data, build)
		handler(build)
	})

	return err
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	b := Builds{}
	go func() {
		b.Watch("http://localhost:8084/events", func(obj interface{}) {
			build := obj.(*api.Build)
			log.Println("Build::::", build)
		})
	}()

	wg.Wait()
}
