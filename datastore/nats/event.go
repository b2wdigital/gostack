package nats

import (
	"context"

	ginats "github.com/b2wdigital/goignite/v2/contrib/nats-io/nats.go.v1"
	"github.com/b2wdigital/gostack/repository"
)

// NewEvent returns a initialized client
func NewEvent(ctx context.Context) repository.Event {
	publisher, _ := ginats.NewDefaultPublisher(ctx)
	return NewClient(publisher)
}
