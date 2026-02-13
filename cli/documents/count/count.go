package count

import (
	"bass-backend/config"
	"bass-backend/database"
	"bass-backend/database/filters"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

const (
	flagBoekjaar             = "boekjaar"
	flagBoekjaarShort        = "j"
	flagBedrijfsnummer       = "bedrijf"
	flagBedrijfsnummerShort  = "b"
	flagDocumentNummber      = "document"
	flagDocumentNummberShort = "d"
)

type command struct {
	cobra.Command
	boekJaar            string
	bedrijfsnummer      string
	documentNummerRange string
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

	result.Flags().StringVarP(&result.boekJaar, flagBoekjaar, flagBoekjaarShort, "", "Boekjaar")
	result.Flags().StringVarP(&result.bedrijfsnummer, flagBedrijfsnummer, flagBedrijfsnummerShort, "", "Bedrijfsnummer")
	result.Flags().StringVarP(&result.documentNummerRange, flagDocumentNummber, flagDocumentNummberShort, "", "Documentnummer interval")

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

	if command.Flags().Changed(flagDocumentNummber) {
		bounds := strings.Split(command.documentNummerRange, "-")

		if len(bounds) != 2 {
			panic("invalid documentnummer range")
		}

		lower := bounds[0]
		upper := bounds[1]

		result = append(result, filters.DocumentNummerBetween(lower, upper))
	}

	return filters.And(result...)
}
