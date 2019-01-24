package cfg

import (
	"github.com/jbvmio/kafkactl/cli/x/out"
	"github.com/spf13/cobra"
)

var list bool

var cmdShowContext = &cobra.Command{
	Use:     "context",
	Aliases: []string{"current", "ctx"},
	Short:   "Display current and available context details",
	Run: func(cmd *cobra.Command, args []string) {
		match := true
		switch match {
		case cmd.CalledAs() == "current":
			out.PrintObject(GetContext(), outFlags.Format)
		case list:
			out.PrintObject(GetContextList(), outFlags.Format)
		default:
			out.PrintObject(GetContext(args...), outFlags.Format)
		}
	},
}

func init() {
	cmdShowContext.Flags().BoolVarP(&list, "list", "l", false, "List available contexts.")
}
