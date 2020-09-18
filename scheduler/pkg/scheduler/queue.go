package scheduler

import (
	"github.com/workhorse/api"
	"sync"
)

type BuildQueue struct {
	queue []*api.Build
	lock  sync.RWMutex
}

func NewBuildQueue() *BuildQueue {
	return &BuildQueue{
		lock: sync.RWMutex{},
	}
}

func (bq *BuildQueue) Add(build *api.Build) {
	bq.lock.Lock()
	defer bq.lock.Unlock()

	bq.queue = append(bq.queue, build)
}

func (bq *BuildQueue) Pop() *api.Build {
	bq.lock.Lock()
	defer bq.lock.Unlock()

	item := bq.queue[len(bq.queue)-1]
	bq.queue[len(bq.queue)-1] = nil

	bq.queue = bq.queue[:len(bq.queue)-1]
	return item
}

func (bq *BuildQueue) Peek() *api.Build {
	item := bq.queue[len(bq.queue)-1]
	return item
}

func (bq *BuildQueue) Length() int {
	return len(bq.queue)
}
