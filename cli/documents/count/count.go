package count

import (
	"bass-backend/config"
	"bass-backend/database"
	"bass-backend/database/filters"
	"fmt"

	"github.com/spf13/cobra"
)

const (
	flagBoekjaar      = "boekjaar"
	FlagBoekjaarShort = "j"
)

type command struct {
	cobra.Command
	boekJaar              int
	bedrijfsnummer        string
	minimumDocumentNummer int
	maximumDocumentNummer int
}

func New() *cobra.Command {
	var result *command

	result = &command{
		Command: cobra.Command{
			Use:   "count",
			Short: "Count documents",
			Args:  cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return result.execute()
			},
		},
	}

	result.Flags().IntVarP(&result.boekJaar, flagBoekjaar, FlagBoekjaarShort, -1, "Boekjaar")

	return &result.Command
}

func (command command) execute() error {
	db, err := database.OpenDatabase(config.DatabasePath)
	if err != nil {
		fmt.Printf("Failed to open database: %s\n", err.Error())
	}
	defer db.Close()

	filter := command.buildFilter()
	count, err := database.CountDocuments(db, filter)
	if err != nil {
		return err
	}

	fmt.Printf("%d\n", count)

	return nil
}

func (command command) buildFilter() filters.Filter {
	result := []filters.Filter{}

	if command.Flags().Changed(flagBoekjaar) {
		result = append(result, filters.Boekjaar(command.boekJaar))
	}

	return filters.And(result...)
}
