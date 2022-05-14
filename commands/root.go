package commands

import (
	"github.com/spf13/cobra"

	"github.com/mimatache/gcurl/commands/call"
	"github.com/mimatache/gcurl/commands/version"
)

func RootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "gcurl",
		Short: "Utlilitary to perform HTTP requests from your CLI",
	}

	version.RegisterTo(root)
	call.RegisterTo(root)

	return root
}
