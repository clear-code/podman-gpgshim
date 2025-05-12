package cmd

import (
	"context"
	"github.com/urfave/cli/v3"
	"os"
)

func Execute() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "show podman-gpgshim version",
	}

	cmd := &cli.Command{
		Name:    "podman-gpgshim (fake GnuPG) for Podman",
		Version: "v1.0.0",
	}

	cmd.Run(context.Background(), os.Args)
}
