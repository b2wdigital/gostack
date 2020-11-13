package cmd

import (
	gsfx "github.com/b2wdigital/gostack/fx"
	"github.com/b2wdigital/gostack/lambda"
	co "github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewLambda() CmdFunc {

	return func(options fx.Option) *co.Command {
		return &co.Command{
			Use:   "lambda",
			Short: "lambda",
			Long:  "",
			Run: func(CmdFunc *co.Command, args []string) {
				gsfx.Run(lambda.HelperModule(options))
			},
		}
	}
}
