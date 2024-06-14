package cli

import "github.com/spf13/cobra"

type App struct {
	Name    string
	Actions map[string]Action
}

func (app *App) toCobraCmd() *cobra.Command {

	rootCmd := &cobra.Command{
		Use: app.Name,
	}

	for name, cmd := range app.Actions {
		rootCmd.AddCommand(cmd.toCobraCmd(name))
	}

	return rootCmd
}

func (app *App) Exec() error {
	return app.toCobraCmd().Execute()
}
