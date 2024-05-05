package root

import (
	"github.com/godunko-mikhail/contexter/cmd/contexter/use_context"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "contexter [help]",
	Short: "Contexter is a tool that helps you manage the local environment of your projects",
	Long:  `Tool with config like kubeconfig but for local environment`,
	Run: func(cmd *cobra.Command, args []string) {
		println("Hello World")
	},
}

func Execute() {
	_ = RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(use_context.UseContext)
}
