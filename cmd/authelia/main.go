package main

import (
	"os"

	"github.com/BlackRoad-OS/authelia/v4/internal/commands"
)

func main() {
	if err := commands.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
