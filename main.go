package main

import (
	"os"

	"github.com/stellaraf/panw-device-state-export/internal/cli"
)

func main() {
	cli.Run(os.Args)
}
