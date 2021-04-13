package nats

import (
	"sync"

	"github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/module/context"
	ginatsfx "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/module/nats-io/nats.go.v1"
	"github.com/b2wdigital/gostack/cloudevents"
	"go.uber.org/fx"
)

var once sync.Once

func HelperModule(extraOptions fx.Option) fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			context.Module(),
			extraOptions,
			ginatsfx.SubscriberModule(),
			cloudevents.HandlerWrapperModule(),
			fx.Provide(
				DefaultOptions,
				NewHelper,
			),
			fx.Invoke(
				func(helper *Helper) {
					helper.Start()
				},
			),
		)
	})

	return options
}
