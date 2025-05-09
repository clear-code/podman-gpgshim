package cmd

import (
    "os"
    "context"

    "github.com/urfave/cli/v3"
)

func Execute() {
    (&cli.Command{}).Run(context.Background(), os.Args)
}
