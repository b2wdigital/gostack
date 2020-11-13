package cmd

import (
	gsfx "github.com/b2wdigital/gostack/fx"
	"github.com/b2wdigital/gostack/nats"
	co "github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewNats() CmdFunc {
	return func(options fx.Option) *co.Command {
		return &co.Command{
			Use:   "nats",
			Short: "nats",
			Long:  "",
			Run: func(CmdFunc *co.Command, args []string) {
				gsfx.Run(nats.HelperModule(options))
			},
		}
	}
}
