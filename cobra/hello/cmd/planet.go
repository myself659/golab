package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var GreetPlanetCmd = &cobra.Command{
	Use: "mars",
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			fmt.Println("About to greet friends from Mars...")
		}
		/* 从root命令继承的选项 */
		lang := cmd.Flag("lang").Value.String()
		fmt.Printf("%s Mars :)\n", greeting(lang))
	},
}
