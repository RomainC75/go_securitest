package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "securitest",
	Long:  "Securitest Api v0.1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Todo API")
	},
}

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "root command",
	Long:  `securytest api`,
	Run: func(cmd *cobra.Command, args []string) {
		// ...
	},
}

func Execute() {
	if err := versionCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
