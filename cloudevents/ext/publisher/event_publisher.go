package publisher

import (
	"context"
	"time"

	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/b2wdigital/gostack/cloudevents"
	"github.com/b2wdigital/gostack/repository"
	"github.com/b2wdigital/gostack/wrapper/provider"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
)

type EventPublisher struct {
	cloudevents.UnimplementedMiddleware
	events repository.Event
}

func NewEventPublisher(events *provider.EventWrapperProvider) cloudevents.Middleware {
	if !IsEnabled() {
		return nil
	}
	return &EventPublisher{events: events}
}

func (p *EventPublisher) AfterAll(ctx context.Context, inouts []*cloudevents.InOut) (context.Context, error) {

	logger := log.FromContext(ctx).WithTypeOf(*p)

	var outs []*v2.Event

	for _, inout := range inouts {
		if inout.Err != nil {
			logger.Warn("the messages could not be published because one or more messages contain errors")
			return ctx, nil
		}
	}

	for _, inout := range inouts {

		out := inout.Out
		in := inout.In

		if out != nil {

			id := uuid.New()

			out.SetID(id.String())
			out.SetExtension("parentId", in.ID())
			out.SetTime(time.Now())

			for key, value := range in.Extensions() {
				out.SetExtension(key, value)
			}

			outs = append(outs, out)
		}

	}

	if er := p.events.Publish(ctx, outs); er != nil {
		return ctx, er
	}

	logger.Info("published events")

	return ctx, nil
}
