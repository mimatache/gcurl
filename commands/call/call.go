package call

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/mimatache/gcurl/request"
)

var (
	ErrOneURLSupported = fmt.Errorf("expected only 1 URL to be given")
)

func RegisterTo(cmd *cobra.Command) {
	call := &cobra.Command{
		Use:   "call",
		Short: "Perform HTTP request. Default action is GET",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("%w; received: %s", ErrOneURLSupported, args)
			}
			client := request.New(args[0])
			response, err := client.Do(http.MethodGet)
			if err != nil {
				return fmt.Errorf("could not perform call to %s; reason: %w", args[0], err)
			}
			fmt.Println(response.GetBody())
			return nil
		},
	}
	cmd.AddCommand(call)
}
