package main

import (
	"os"

	"github.com/AmirAbaskohi/AmirhosseinChain/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
