package list

import (
	"bass-backend/database"
	"fmt"

	"github.com/spf13/cobra"
)

type command struct {
	cobra.Command
}

func New() *cobra.Command {
	var result *command

	result = &command{
		Command: cobra.Command{
			Use:   "list",
			Short: "List documents",
			Args:  cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return result.execute()
			},
		},
	}

	return &result.Command
}

func (command command) execute() error {
	path := "bookkeeping.db"

	db, err := database.CreateDatabase(path)
	if err != nil {
		fmt.Printf("Failed to create database: %s\n", err.Error())
	}
	db.Close()

	return nil
}
