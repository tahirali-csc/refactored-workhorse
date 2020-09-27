package logstorage

import (
	"context"
)

type LogStore interface {
	// Find returns a log stream from the datastore.
	//Find(ctx context.Context, stage int64) (io.ReadCloser, error)

	// Create writes copies the log stream from Reader r to the datastore.
	Write(ctx context.Context, stepId int64, line []byte) error
}
