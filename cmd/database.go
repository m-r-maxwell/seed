package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "NOT IMPLEMENTED YET Database commands for managing the database",
	Long:  "Database commands for managing the database, including migrations and seeding.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Database not implemented yet.")
		// TODO: Implement database command
		// would want to add a commmand for building
		// create: sqlite, postgres, mysql
		// migrate: run migrations
		// seed: generate data by calling pkg.GenerateData functions

	},
}

func init() {
	rootCmd.AddCommand(databaseCmd)
}
