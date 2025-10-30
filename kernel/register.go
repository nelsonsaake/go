package kernel

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

type Handler func([]string) error

var registry = map[string]Handler{}

func Register(name string, handler Handler) {
	registry[name] = handler
}

func register(c *cobra.Command, name string, h Handler) {
	c.AddCommand(&cobra.Command{
		Use: name,
		Run: func(cmd *cobra.Command, args []string) {

			if IsVerbose() {
				fmt.Println(name, "...")
			}

			err := h(args)
			if err != nil {
				log.Fatalf("error executing %v: %v", name, err)
			}
		},
	})
}

func loadRegister(c *cobra.Command) {
	for name, handler := range registry {
		register(c, name, handler)
	}
}
