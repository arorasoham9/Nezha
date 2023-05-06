package cmd

import (
	b "Nezha-cli/basic"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Nezha running on your host machine",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if !cmd.Flags().Changed("HP") {
			fmt.Println("Could not resolve 'Host:Port'")
		}else{
			b.InitRest()
			var loggedUser b.LoggedUser
			loggedUser, check := b.LoginStatus(); if check {
				fmt.Println("Already logged in as:", loggedUser.USERNAME, "Last login:", loggedUser.LOGGED_AT )
			}else{
				var newUserLoginRequest b.LoginUser
				newUserLoginRequest.USERNAME, newUserLoginRequest.PASSWORD = b.GetCredentials()
				newUserLoginRequest.CREATED_AT = time.Now().GoString()
				HOSTPORT, _ := cmd.Flags().GetString("HP")
				newUserLoginResponse, check := b.SendLoginRequest(newUserLoginRequest, HOSTPORT)
				if !check{
					fmt.Println(newUserLoginResponse.MESSAGE)
				}
				fmt.Println("Login success! Your login details are stored unencrypted locally.")
		} 
		}
	},
}


func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.PersistentFlags().String("HP", "", "Host:Port") 
}
