package cmd

import (
	"os"

	"github.com/b2wdigital/goignite/v2/contrib/spf13/cobra.v1"
	co "github.com/spf13/cobra"
	"go.uber.org/fx"
)

type CmdFunc func(fx.Option) *co.Command

func Run(options fx.Option, c ...CmdFunc) error {

	var cmds []*co.Command

	for _, v := range c {
		cmds = append(cmds, v(options))
	}

	rootCmd := &co.Command{
		Use:   "gostack",
		Short: "gostack",
		Long:  "",
	}

	if def := os.Getenv("DEFAULT_CMD"); def != "" {
		cmd, _, err := rootCmd.Find(os.Args[1:])
		if err == nil && cmd.Use == rootCmd.Use {
			args := append([]string{def}, os.Args[1:]...)
			rootCmd.SetArgs(args)
		}
	}

	return cobra.Run(rootCmd, cmds...)
}
