package database

import (
	"bass-backend/cli/database/create"
	importdata "bass-backend/cli/database/import"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	result := cobra.Command{
		Use:   "database",
		Short: "Database commands",
	}

	result.AddCommand(create.New())
	result.AddCommand(importdata.New())

	return &result
}
