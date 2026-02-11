package database

import (
	"bass-backend/cli/database/create"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	result := cobra.Command{
		Use:   "database",
		Short: "Database commands",
	}

	result.AddCommand(create.New())

	return &result
}
