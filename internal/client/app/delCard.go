package app

import (
	"log"

	"github.com/spf13/cobra"
)

var delCard = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "delcard",
	Short: "Delete user card by id",
	Long: `
This command remove card
Usage: delcard -i \"card_id\" 
Flags:
  -i, --id string Card id
  -p, --password string   User password value.`,
	Run: func(cmd *cobra.Command, args []string) {
		clientUseCase.DelCard(userPassword, delCardID)
	},
}

var delCardID string //nolint:gochecknoglobals // cobra style guide

func init() {
	rootCmd.AddCommand(delCard)
	delCard.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	delCard.Flags().StringVarP(&delCardID, "id", "i", "", "Card id")

	if err := delCard.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := delCard.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
