package root

import (
	"bass-backend/cli/database"
	"bass-backend/cli/documents"
	"bass-backend/cli/server"
	"log/slog"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	verbose := false

	result := cobra.Command{
		Use:   "backend",
		Short: "Command line interface",
		Long:  "Command line interface",
	}

	cobra.OnInitialize(func() {
		if verbose {
			slog.SetLogLoggerLevel(slog.LevelDebug)
			slog.Info("Verbose mode enabled")
		}
	})

	result.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")

	result.AddCommand(database.New())
	result.AddCommand(server.New())
	result.AddCommand(documents.New())

	return &result
}
