/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/bhoriuchi/go-bunyan/pkg/bunyan"
)

var version, strict, color bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-bunyan",
	Short: "Filter and pretty-print Bunyan log file content",
	Long: `bunyan [OPTIONS] [FILE ...]
	... | bunyan [OPTIONS]

	bunyan [OPTIONS] -p PID`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		parseArgs := bunyan.ParsedArgs{}
		for _, arg := range args {
			switch arg {
			case "--version":
				version = true
			case "--strict":
				parseArgs.Strict = true
			case "--color":
				parseArgs.Color = true
			case "--no-color":
				parseArgs.Color = false
			case "--pager":
				parseArgs.Pagination = true
			case "--no-pager":
				parseArgs.Pagination = false
				// case "--output":

			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "While executing go-bunyan %v", err)
		time.Sleep(10 * time.Millisecond)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-bunyan.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVar(&version, "version", false, "print version of this command and exit")
	rootCmd.Flags().BoolVar(&strict, "strict", false, `Suppress all but legal Bunyan JSON log lines. 
	By default non-JSON, and non-Bunyan lines are passed through.`)
	rootCmd.Flags().Bool("color", false, "Colorize output. Defaults to try if output, stream is a TTY.")
	rootCmd.Flags().Bool("no-color", false, "Force no coloring (e.g. terminal doesn't support it)")
	rootCmd.Flags().String("time-format", "utc", "Display time field in local time or UTC.")
	rootCmd.Flags().String("output", "long", `Specify an output mode/format. One of
	long: (the default) pretty
	json: JSON output, 2-space indent
	json-N: JSON output, N-space indent, e.g. "json-4"
	bunyan: 0 indented JSON, bunyan's native format'
	short: like "long", but more concise
	simple: level, followed by "-" and then the message`)
	rootCmd.Flags().Bool("pager", false, `Pipe output into 'less' (or $PAGER if set), if 
	stdout is a TTY. This overrides $BUNYAN_NO_PAGER.`)
	rootCmd.Flags().Bool("no-pager", false, "Do not pipe output into a pager.")

}
