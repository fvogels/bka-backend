package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	result := cobra.Command{
		Use:   "create",
		Short: "Create empty database",
		RunE:  execute,
	}

	return &result
}

func execute(cmd *cobra.Command, args []string) error {
	fmt.Println("Creating database!")

	return nil
}
