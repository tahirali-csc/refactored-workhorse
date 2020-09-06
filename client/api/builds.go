package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

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

func (b *Builds) BindToNode(binding api.BuildNodeBinding) {
	client := http.Client{}

	dt, _ := json.Marshal(binding)

	res, err := client.Post("http://localhost:8081/buildbinding", "application/json", bytes.NewReader(dt))
	log.Println(res)
	if err != nil {
		log.Println("Error", err)
	}

}
