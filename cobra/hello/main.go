package main

import (
	"fmt"
	"os"

	"github.com/myself659/golab/cobra/hello/cmd"
)

func main() {
	if err := cmd.RootCmdEx.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
