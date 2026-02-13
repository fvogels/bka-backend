package list

import (
	"bass-backend/config"
	"bass-backend/database"
	"bass-backend/database/queries"
	"bass-backend/model"
	"encoding/json"
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
			Use:   "list",
			Short: "List documents",
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

	query := command.buildQuery()
	documents, err := query.Execute(db)
	if err != nil {
		return err
	}

	formattedDocuments, err := json.MarshalIndent(documents, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to convert documents to JSON format: %w", err)
	}
	fmt.Println(string(formattedDocuments))

	return nil
}

func (command command) buildQuery() *queries.ListDocumentsQuery {
	query := queries.ListDocuments()

	if command.Flags().Changed(flagBoekjaar) {
		query.WithBoekjaar(model.NewBoekJaar(command.boekJaar))
	}

	if command.Flags().Changed(flagBedrijfsnummer) {
		query.WithBedrijfsnummer(model.NewBedrijfsnummer(command.bedrijfsnummer))
	}

	if command.Flags().Changed(flagDocumentNummber) {
		bounds := strings.Split(command.documentNummerRange, "-")

		if len(bounds) != 2 {
			panic("invalid documentnummer range")
		}

		lower := bounds[0]
		upper := bounds[1]

		if len(lower) >= 1 && lower[0] == '>' {
			lower = fmt.Sprintf("%010s", lower[1:])
		}

		if len(upper) >= 1 && upper[0] == '>' {
			upper = fmt.Sprintf("%010s", upper[1:])
		}

		query.WithDocumentNummerBetween(lower, upper)
	}

	return query
}
