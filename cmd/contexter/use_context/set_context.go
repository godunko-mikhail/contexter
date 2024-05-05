package use_context

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var UseContext = &cobra.Command{
	Use: "use-context",
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		if err != nil {
		}
		fmt.Println(home)
	},
}
