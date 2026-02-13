package documents

import (
	"bass-backend/cli/documents/count"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	result := cobra.Command{
		Use:   "documents",
		Short: "Document commands",
	}

	result.AddCommand(count.New())

	return &result
}
