package count

import (
	"bass-backend/config"
	"bass-backend/database"
	"bass-backend/database/filters"
	"fmt"

	"github.com/spf13/cobra"
)

const (
	flagBoekjaar            = "boekjaar"
	flagBoekjaarShort       = "j"
	flagBedrijfsnummer      = "bedrijf"
	flagBedrijfsnummerShort = "b"
)

type command struct {
	cobra.Command
	boekJaar              string
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

	result.Flags().StringVarP(&result.boekJaar, flagBoekjaar, flagBoekjaarShort, "0000", "Boekjaar")
	result.Flags().StringVarP(&result.bedrijfsnummer, flagBedrijfsnummer, flagBedrijfsnummerShort, "0000", "Bedrijfsnummer")

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

	if command.Flags().Changed(flagBedrijfsnummer) {
		result = append(result, filters.Bedrijfsnummer(command.bedrijfsnummer))
	}

	return filters.And(result...)
}
