//MIT License
//
//Copyright (c) 2018 ChenQuan
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

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
