package lambda

import (
	"sync"

	contextfx "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/module/context"
	"github.com/b2wdigital/gostack/cloudevents"
	"go.uber.org/fx"
)

var once sync.Once

func HelperModule(extraOptions fx.Option) fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			extraOptions,
			cloudevents.HandlerWrapperModule(),
			fx.Provide(
				NewDefaultHelper,
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
