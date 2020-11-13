package fx

import (
	gifx "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1"
	"go.uber.org/fx"
)

func Run(options fx.Option) {
	gifx.New(options).Run()
}
