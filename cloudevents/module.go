package cloudevents

import (
	"sync"

	cloudevents "github.com/b2wdigital/goignite/v2/contrib/cloudevents/sdk-go.v2"
	contextfx "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/module/context"
	"go.uber.org/fx"
)

var handlerWrapperOnce sync.Once

func HandlerWrapperModule() fx.Option {
	options := fx.Options()

	handlerWrapperOnce.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				DefaultHandlerWrapperOptions,
				func(handler cloudevents.Handler, options *HandlerWrapperOptions, m []Middleware) *HandlerWrapper {
					return NewHandlerWrapper(handler, options, m...)
				},
			),
		)
	})

	return options
}

var helperOnce sync.Once

func HelperModule(extraOptions fx.Option) fx.Option {
	options := fx.Options()

	helperOnce.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			extraOptions,
			HandlerWrapperModule(),
			fx.Provide(
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
