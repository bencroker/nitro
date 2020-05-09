package cmd

import (
	"github.com/spf13/cobra"

	"github.com/craftcms/nitro/config"
	"github.com/craftcms/nitro/internal/nitro"
)

var xdebugConfigureCommand = &cobra.Command{
	Use:    "configure",
	Short:  "Configure Xdebug on a machine",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		machine := flagMachineName

		php := config.GetString("php", flagPhpVersion)

		var actions []nitro.Action
		xdebugConfigureAction, err := nitro.ConfigureXdebug(machine, php)
		if err != nil {
			return err
		}
		actions = append(actions, *xdebugConfigureAction)

		restartPhpFpmAction, err := nitro.RestartPhpFpm(machine, php)
		if err != nil {
			return err
		}
		actions = append(actions, *restartPhpFpmAction)

		return nitro.Run(nitro.NewMultipassRunner("multipass"), actions)
	},
}
