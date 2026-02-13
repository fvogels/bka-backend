package importdata

import (
	"bass-backend/config"
	"bass-backend/database"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type importCommand struct {
	cobra.Command
	documentPath string
	segmentPath  string
}

func New() *cobra.Command {
	var result *importCommand

	result = &importCommand{
		Command: cobra.Command{
			Use:   "import",
			Short: "Import CSV data",
			RunE: func(cmd *cobra.Command, args []string) error {
				return result.execute()
			},
			Args: cobra.NoArgs,
		},
	}

	result.Flags().StringVar(&result.documentPath, "documents", "", "Documents CSV")
	result.Flags().StringVar(&result.segmentPath, "segments", "", "Segments CSV")

	if err := result.MarkFlagRequired("documents"); err != nil {
		panic("failed to mark documents flag as required")
	}

	if err := result.MarkFlagRequired("segments"); err != nil {
		panic("failed to mark segments flag as required")
	}

	return &result.Command
}

func (command importCommand) execute() error {
	path := config.DatabasePath

	db, err := database.OpenDatabase(path)
	if err != nil {
		return fmt.Errorf("Failed to create database: %s\n", err.Error())
	}
	defer db.Close()

	documentData, err := os.Open(command.documentPath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", command.documentPath, err)
	}
	defer documentData.Close()

	segmentData, err := os.Open(command.segmentPath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", command.segmentPath, err)
	}
	defer segmentData.Close()

	if err := database.ImportData(db, documentData, segmentData); err != nil {
		return fmt.Errorf("failed to import data: %w", err)
	}

	return nil
}
