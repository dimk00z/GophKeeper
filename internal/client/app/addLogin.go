package app

import (
	"log"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/spf13/cobra"
)

var addLogin = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "addlogin",
	Short: "Add login",
	Long: `
This command add logit for site
Usage: addlogin -p \"user_password\" 
Flags:
  -h, --help              help for addlogin
  -l, --login string      Site login
  -p, --password string   User password value.
  -s, --secret string     Site password|secret
  -t, --title string      Login title
  -u, --uri string        Site endloint`,
	Run: func(cmd *cobra.Command, args []string) {
		clientUseCase.AddLogin(userPassword, &loginForAdditing)
	},
}

var loginForAdditing entity.Login //nolint:gochecknoglobals // cobra style guide

func init() {
	rootCmd.AddCommand(addLogin)
	addLogin.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")

	addLogin.Flags().StringVarP(&loginForAdditing.Name, "title", "t", "", "Login title")
	addLogin.Flags().StringVarP(&loginForAdditing.Login, "login", "l", "", "Site login")
	addLogin.Flags().StringVarP(&loginForAdditing.Password, "secret", "s", "", "Site password|secret")
	addLogin.Flags().StringVarP(&loginForAdditing.URI, "uri", "u", "", "Site endloint")

	if err := addLogin.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := addLogin.MarkFlagRequired("title"); err != nil {
		log.Fatal(err)
	}
}
