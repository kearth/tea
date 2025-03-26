package container

import "context"

type Component interface {
	Object
	Init(ctx context.Context) error
}
