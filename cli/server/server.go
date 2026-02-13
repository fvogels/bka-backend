package server

import (
	"bass-backend/cli/server/run"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	result := cobra.Command{
		Use:   "server",
		Short: "Server commands",
	}

	result.AddCommand(run.New())

	return &result
}
