package cfg

import (
	"strings"

	"github.com/jbvmio/kafkactl/cli/x/out"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdUseContext = &cobra.Command{
	Use:     "context",
	Aliases: []string{"ctx"},
	Short:   "switch active context",
	Long:    `command to switch active context`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		context := strings.Join(args, " ")
		ctx := GetContext(context)
		viper.Set("current-context", ctx.Name)
		if err := viper.WriteConfig(); err != nil {
			out.Failf("Unable to write config: %v", err)
		}
		out.Infof("Context changed to %v", ctx.Name)
	},
}

func init() {

}
