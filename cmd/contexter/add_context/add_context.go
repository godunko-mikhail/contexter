package add_context

import (
	"fmt"
	"github.com/godunko-mikhail/contexter/pkg/config"
	"github.com/godunko-mikhail/contexter/pkg/contexter"
	"github.com/spf13/cobra"
)

var (
	context = &config.Context{}
)

var Cmd = &cobra.Command{
	Use:   "add-context [context]",
	Short: "Adds a new context to the contexter config",

	RunE: func(cmd *cobra.Command, args []string) error {
		contexter.Init("")

		if !config.Exists() {
			fmt.Println("No config file found, please run 'contexter init'")
			return nil
		}

		cfg, err := config.Load()
		if err != nil {
			return err
		}
		if _, err := cfg.GetContext("t"); err == nil {
			fmt.Println("Context with this name already exists")
			return nil
		}

		err = cfg.AddContext(context)
		if err != nil {
			return err
		}

		err = config.Write(cfg)
		if err != nil {
			return err
		}

		fmt.Println("Added context successfully")
		return nil
	},
}

func init() {
	Cmd.Flags().StringVarP(&context.Name, "name", "n", "", "The name of the context to add")
	Cmd.Flags().StringVarP(&context.Executable, "executable", "x", "", "Executable for new environment")
	Cmd.Flags().StringArrayVarP(&context.Commands, "commands", "c", []string{}, "Commands in new environment")
	Cmd.Flags().StringToStringVarP(&context.Envs, "envs", "e", map[string]string{}, "Environment variables")
	_ = Cmd.MarkFlagRequired("name")
	_ = Cmd.MarkFlagRequired("executable")
}
