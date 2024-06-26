package use_context

import (
	"fmt"
	"github.com/godunko-mikhail/contexter/pkg/config"
	"github.com/godunko-mikhail/contexter/pkg/contexter"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "use-context [context]",
	Short: "Starts new process with defined context",
	RunE: func(cmd *cobra.Command, args []string) error {
		contexter.Init("")

		if !config.Exists() {
			fmt.Printf("No config file found, please run 'contexter init'\n")
			return nil
		}

		cfg, err := config.Load()
		if err != nil {
			return err
		}
		print("%v", cfg)

		return nil
	},
}
