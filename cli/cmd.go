package cli

import "github.com/spf13/cobra"

type Action struct {
	Run     RunFunc
	Actions map[string]Action
}

func (c *Action) toCobraCmd(name string) *cobra.Command {

	rootCmd := &cobra.Command{
		Use: name,
		Run: func(cmd *cobra.Command, args []string) {
			c.Run(args...)
		},
	}

	for name, cmd := range c.Actions {
		rootCmd.AddCommand(cmd.toCobraCmd(name))
	}

	return rootCmd
}
