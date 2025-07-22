package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const asciiArt = `
 _____               _ 
/  ___|             | |
\ '--.  ___  ___  __| |
 '--. \/ _ \/ _ \/ _' |
/\__/ /  __/  __/ (_| |
\____/ \___|\___|\__,_|
                       

`

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "seed",
	Version: version,
	Short:   "Seed is a CLI tool for generate data and databases",
	Long:    "Seed is a CLI tool for generating data and databases. It provides various commands to help you create and manage your data efficiently.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(asciiArt)
		cmd.Help()
	},
}

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "Help for seed command")
	rootCmd.Flags().BoolP("version", "v", false, "Print the version number")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
