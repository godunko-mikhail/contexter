package initialize

import (
	"fmt"
	"github.com/godunko-mikhail/contexter/pkg/config"
	"github.com/godunko-mikhail/contexter/pkg/contexter"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize contexter config",
	RunE: func(cmd *cobra.Command, args []string) error {
		contexter.Init("")

		if config.Exists() {
			fmt.Println("Already initialized")
			return nil
		}

		err := config.Create()
		if err != nil {
			return err
		}

		fmt.Println("Successfully initialized")
		return nil
	},
}
