package run

import (
	"bass-backend/rest"

	"github.com/spf13/cobra"
)

type command struct {
	cobra.Command
}

func New() *cobra.Command {
	var result *command

	result = &command{
		Command: cobra.Command{
			Use:   "run",
			Short: "Run server",
			Args:  cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return result.execute()
			},
		},
	}

	return &result.Command
}

func (command command) execute() error {
	return rest.StartServer()
}
