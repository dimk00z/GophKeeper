package app

import (
	"log"

	"github.com/spf13/cobra"
)

var delLogin = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "dellogin",
	Short: "Delete user login by id",
	Long: `
This command remove login
Usage: delcard -i \"login_id\" 
Flags:
  -i, --id string Card id
  -p, --password string   User password value.`,
	Run: func(cmd *cobra.Command, args []string) {
		clientUseCase.DelLogin(userPassword, delLoginID)
	},
}

var delLoginID string //nolint:gochecknoglobals // cobra style guide

func init() {
	rootCmd.AddCommand(delLogin)
	delLogin.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	delLogin.Flags().StringVarP(&delLoginID, "id", "i", "", "Card id")

	if err := delLogin.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := delLogin.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
