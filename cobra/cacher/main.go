package main

import (
	"fmt"
	"os"

	"github.com/myself659/golab/cobra/cacher/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
