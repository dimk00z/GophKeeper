package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

type loginUser struct {
	login    string
	password string
}

var RequiredUserArgs = 2 //nolint:gochecknoglobals // cobra style guide

var loginUserCmd = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "login",
	Short: "Login user to the service",
	Long: `
This command login user.
Usage: gophkeeperclient login user_login user_password`,
	Args: cobra.MinimumNArgs(RequiredUserArgs),
	Run: func(cmd *cobra.Command, args []string) {
		login := loginUser{
			login:    args[0],
			password: args[1],
		}
		fmt.Println(login)
		// TODO: add login logic
	},
}

func init() {
	rootCmd.AddCommand(loginUserCmd)
}
