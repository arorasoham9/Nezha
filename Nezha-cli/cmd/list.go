/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	// "net/http"
	// "strings"

	"github.com/spf13/cobra"
)

const listLogURL = "/list/"

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List deployed applications.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}


func init() {
	rootCmd.AddCommand(listCmd)
}
