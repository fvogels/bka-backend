package documents

import (
	"bass-backend/cli/documents/count"
	"bass-backend/cli/documents/list"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	result := cobra.Command{
		Use:   "documents",
		Short: "Document commands",
	}

	result.AddCommand(count.New())
	result.AddCommand(list.New())

	return &result
}
