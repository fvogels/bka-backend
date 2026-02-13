package create

import (
	"bass-backend/database"
	"fmt"

	"github.com/spf13/cobra"
)

type createCommand struct {
	cobra.Command
}

func New() *cobra.Command {
	var result *createCommand

	result = &createCommand{
		Command: cobra.Command{
			Use:   "create",
			Short: "Create empty database",
			Args:  cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return result.execute()
			},
		},
	}

	return &result.Command
}

func (command createCommand) execute() error {
	path := "bookkeeping.db"

	db, err := database.CreateDatabase(path)
	if err != nil {
		fmt.Printf("Failed to create database: %s\n", err.Error())
	}
	db.Close()

	return nil
}
