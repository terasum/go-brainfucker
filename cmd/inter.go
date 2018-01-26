package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/terasum/gobf/interpreter"
	"io/ioutil"
	"os"
)

var (
	verbose bool
)

func init() {
	RootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

//RootCmd is the root command of gobf
var RootCmd = &cobra.Command{
	Use:   "gobf",
	Short: "BrainFuck the language interpreter in go ",
	Long: `gobf is a command line interpreter for BrainFuck language.
you can use:
    gobf source.bf
to interpreter the BrainFuck code
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("missing code file path")
			os.Exit(1)
		}
		code, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Printf("cannot read the code %s\n", err.Error())
			os.Exit(1)
		}
		inter := interpreter.NewInterpreter(code, verbose)
		ret, err := inter.Run()
		if err != nil {
			fmt.Printf("excute filed %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Println(ret)
	},
}
