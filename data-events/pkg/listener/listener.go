package listener

import (
	"database/sql"
	"log"
	"time"

	"github.com/lib/pq"
)

type DatabaseEventsListener struct {
	EventChannel chan string
}

func New() *DatabaseEventsListener {
	return &DatabaseEventsListener{
		EventChannel: make(chan string),
	}
}

func (listener *DatabaseEventsListener) Start() {
	conninfo := "dbname=workflow user=dev password=dev host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conninfo)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	eventListener := pq.NewListener(conninfo, 10*time.Second, time.Minute, nil)
	err = eventListener.Listen("events")
	if err != nil {
		panic(err)
	}

	listener.waitForNotification(eventListener)
}

func (listener *DatabaseEventsListener) waitForNotification(l *pq.Listener) {
	for {
		select {

		case e := <-l.Notify:

			//Send event async. TODO: Will review
			select {
			case listener.EventChannel <- e.Extra:
			default:
			}

		case <-time.After(90 * time.Second):
			log.Println("Received no events for 90 seconds, checking connection")
			go func() {
				l.Ping()
			}()
		}
	}
}
