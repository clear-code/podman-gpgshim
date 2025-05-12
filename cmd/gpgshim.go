package cmd

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"os"
)

func Execute() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "show podman-gpgshim version",
	}

	var recipient string
	var output string
	var homedir string
	var passphrase_fd string
	var export_key string
	var export_secret_key string
	cmd := &cli.Command{
		Name:    "podman-gpgshim (fake GnuPG) for Podman",
		Version: "v1.0.0",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "encrypt",
				Usage:   "encrypt data",
				Aliases: []string{"e"},
			},
			&cli.BoolFlag{
				Name:    "decrypt",
				Usage:   "decrypt data",
				Aliases: []string{"d"},
			},
			&cli.BoolFlag{
				Name:  "batch",
				Usage: "execute in batch mode",
			},
			&cli.BoolFlag{
				Name:  "with-colons",
				Usage: "specify output format for machine-readable",
			},
			&cli.StringFlag{
				Name:        "recipient",
				Usage:       "encrypt for `USER-ID`",
				Aliases:     []string{"r"},
				Destination: &recipient,
			},
			&cli.StringFlag{
				Name:        "output",
				Usage:       "write output to `FILE`",
				Aliases:     []string{"o"},
				Destination: &output,
			},
			&cli.StringFlag{
				Name:        "homedir",
				Usage:       "specify home directory to `HOME`",
				Destination: &homedir,
			},
			&cli.StringFlag{
				Name:        "passphrase-fd",
				Usage:       "specify `FD` which was used to pass passphrase to",
				Destination: &passphrase_fd,
			},
			&cli.StringFlag{
				Name:        "export-secret-key",
				Usage:       "specify secret `KEY` to be exported",
				Destination: &export_secret_key,
			},
			&cli.StringFlag{
				Name:        "export",
				Usage:       "specify `KEY` to be exported",
				Destination: &export_key,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println(cmd.String("recipient"))
			return nil
		},
	}

	cmd.Run(context.Background(), os.Args)
}
