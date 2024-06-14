package cli

import "github.com/spf13/cobra"

var app = App{
	Name: "npanel",
	Actions: map[string]Action{
		"ls": {
			Run: printArgs,
		},
		"ls2": {
			Run: printArgs,
		},
	},
}

func Main() {
	cobra.CheckErr(app.exec())
}
