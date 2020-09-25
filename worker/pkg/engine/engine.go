package engine

import (
	"context"
	"io"
)

type Engine interface {
	Start(ctx context.Context, spec *Spec, step *Step) error
	Create(ctx context.Context, spec *Spec, step *Step) error
	Tail(context.Context, *Spec, *Step) (io.ReadCloser, error)
}
