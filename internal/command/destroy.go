package command

import (
	"errors"

	"github.com/urfave/cli/v2"
)

// Destroy will completely destroy a machine
func Destroy() *cli.Command {
	return &cli.Command{
		Name:        "destroy",
		Usage:       "Destroy machine",
		Description: "By default, when deleting a machine it is soft-deleted which means it can be recovered. This command will destroy the machine making it unrecoverable.",
		Action:      destroyAction,
	}
}

func destroyAction(c *cli.Context) error {
	return errors.New("not implemented")
}