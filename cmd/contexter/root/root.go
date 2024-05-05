package root

import (
	"github.com/godunko-mikhail/contexter/cmd/contexter/add_context"
	"github.com/godunko-mikhail/contexter/cmd/contexter/initialize"
	"github.com/godunko-mikhail/contexter/cmd/contexter/use_context"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "contexter [help]",
	Short: "Contexter is a tool that helps you manage the local environment of your projects",
	Long:  "Tool with config like kubeconfig but for local environment",
	Run: func(cmd *cobra.Command, args []string) {
		println("Hello World")
	},
}

func Execute() {
	_ = cmd.Execute()
}

func init() {
	cmd.AddCommand(use_context.Cmd)
	cmd.AddCommand(initialize.Cmd)
	cmd.AddCommand(add_context.Cmd)
}
