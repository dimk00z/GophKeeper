package app

import (
	"log"

	"github.com/spf13/cobra"
)

var getLogin = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "getlogin",
	Short: "Show user login by id",
	Long: `
This command getlogin
Usage: getlogin -i \"login_id\" 
Flags:
  -i, --id string Login id
  -p, --password string   User password value.`,
	Run: func(cmd *cobra.Command, args []string) {
		clientUseCase.ShowLogin(userPassword, getLoginID)
	},
}

var getLoginID string //nolint:gochecknoglobals // cobra style guide

func init() {
	rootCmd.AddCommand(getLogin)
	getLogin.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	getLogin.Flags().StringVarP(&getLoginID, "id", "i", "", "Card id")

	if err := getLogin.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := getLogin.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
