package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Verbose bool

/*
init函数 在main函数执行前执行，且有相应的顺序
*/
func init() {
	/* 标记flag，表示是与否 */
	RootCmdEx.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	/* 选项 应用示例 --lang cn */
	RootCmdEx.PersistentFlags().String("lang", "en", "language to use")

	/* 添加子命令 */
	RootCmdEx.AddCommand(GreetPlanetCmd)
	/* 添加子命令 */
	RootCmdEx.AddCommand(GreetTravellerCmd)
}

/*
定义根命令
*/
var RootCmdEx = &cobra.Command{
	Use: "hello",
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			fmt.Println("About to greet friends from Earth...")
		}
		lang := cmd.Flag("lang").Value.String()
		fmt.Println("lang:", lang)
		fmt.Printf("%s world :)\n", greeting(lang))
	},
}
