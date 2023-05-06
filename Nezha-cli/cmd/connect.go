package cmd

import (
	"github.com/spf13/cobra"
)


// listCmd represents the list command
var connectCmd = &cobra.Command{
	Use:   "list",
	Short: "Connect to deployed applications.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}


func init() {
	rootCmd.AddCommand(listCmd)
}
