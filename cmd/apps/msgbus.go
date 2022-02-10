package apps

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kubegems/gems/pkg/msgbus/options"
	"github.com/kubegems/gems/pkg/msgbus"
	"github.com/kubegems/gems/pkg/utils/config"
	"github.com/kubegems/gems/pkg/version"
	"github.com/spf13/cobra"
)

func NewMsgbusCmd() *cobra.Command {
	options := options.DefaultOptions()
	cmd := &cobra.Command{
		Use:          "msgbus",
		Short:        "run msgbus",
		SilenceUsage: true,
		Version:      version.Get().String(),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := config.Parse(cmd.Flags()); err != nil {
				return err
			}
			ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
			defer cancel()

			return msgbus.Run(ctx, options)
		},
	}
	cmd.AddCommand(genCfgCmd)
	options.RegistFlags("", cmd.Flags())
	return cmd
}