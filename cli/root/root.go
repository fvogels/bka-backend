package root

import (
	"bass-backend/cli/database"
	"bass-backend/cli/server"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	result := cobra.Command{
		Use:   "backend",
		Short: "Command line interface",
		Long:  "Command line interface",
	}

	result.AddCommand(database.New())
	result.AddCommand(server.New())

	return &result
}
