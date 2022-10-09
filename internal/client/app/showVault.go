package app

import (
	"log"

	"github.com/spf13/cobra"
)

var showVault = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "showvault",
	Short: "Show user vault",
	Long: `
This command show user vault
Usage: showvault -o \"a|c|l|n\" 
Flags:
  -o, --option string     Option for listing (default "a")
	a - all
	c - cards
	l - logins
	n - notes
	b - bynaries
  -p, --password string   User password value.`,
	Run: func(cmd *cobra.Command, args []string) {
		clientUseCase.ShowVault(userPassword, showVaultOption)
	},
}

var showVaultOption string //nolint:gochecknoglobals // cobra style guide

func init() {
	rootCmd.AddCommand(showVault)
	showVault.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	showVault.Flags().StringVarP(&showVaultOption, "option", "o", "a", "Option for listing")

	if err := showVault.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
}
