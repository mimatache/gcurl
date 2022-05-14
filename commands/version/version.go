package version

import (
	"fmt"

	"github.com/mimatache/gcurl/version"
	"github.com/spf13/cobra"
)

func RegisterTo(cmd *cobra.Command) {
	version := &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(version.AppInfo())
			return nil
		},
	}
	cmd.AddCommand(version)
}
