package main

import (
	"net/http"

	"gihtub.com/workhorse/data-events/pkg/eventmapper"
	"gihtub.com/workhorse/data-events/pkg/listener"
)

func main() {
	//database events listener
	dbEventsListener := listener.New()

	//events to SSE stream mapper
	eventsMapper := eventmapper.New(dbEventsListener)

	mux := http.NewServeMux()
	mux.HandleFunc("/events", eventsMapper.HTTPHandler)

	//Start database events listener
	go dbEventsListener.Start()

	//Watch for database events and send to SSE events
	go eventsMapper.WatchEvents()

	http.ListenAndServe(":8084", mux)
}
