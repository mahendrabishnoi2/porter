package main

import (
	"get.porter.sh/porter/pkg/porter"
	"github.com/spf13/cobra"
)

func buildHelloCommand(p *porter.Porter) *cobra.Command {
	// Store arguments and flags specified by the user
	opts := porter.HelloOptions{}

	// Define the `porter hello` command
	cmd := &cobra.Command{
		Use:   "hello",
		Short: "Say hello",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.Validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return p.Hello(opts)
		},
	}

	// Define the --name flag. Allow using -n too.
	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Your name")
	return cmd
}
