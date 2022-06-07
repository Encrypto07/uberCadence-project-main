package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "uberCadence-project",
	Short: "Run basic predefined example workflows",
	Long: `Using this executable, multiple examples related to Uber Cadence, can run.
These examples involve various use cases of workflows on how they can be implemented in practical environments.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
