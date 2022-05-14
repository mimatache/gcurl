package main

import (
	"fmt"
	"os"

	"github.com/mimatache/gcurl/commands"
)

func main() {
	if err := commands.RootCommand().Execute(); err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}
}
