package cli

import (
	"bass-backend/cli/root"
	"log/slog"
	"os"
)

func ProcessCommandLineArguments() {
	rootCommand := root.New()

	if err := rootCommand.Execute(); err != nil {
		slog.Debug(
			"Error while processing command line arguments",
			slog.String("error", err.Error()),
		)

		os.Exit(1)
	}
}
