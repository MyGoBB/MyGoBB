package main

import (
	"fmt"
	"os"

	"github.com/MyGoBB/MyGoBB/command"
)

func main() {
	if err := command.RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
