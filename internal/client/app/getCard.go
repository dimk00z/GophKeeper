package app

import (
	"log"

	"github.com/spf13/cobra"
)

var getCard = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "getcard",
	Short: "Show user card by id",
	Long: `
This command add card
Usage: showvault -c \"card_id\" 
Flags:
  -c, --card string Card id
  -p, --password string   User password value.`,
	Run: func(cmd *cobra.Command, args []string) {
		clientUseCase.ShowCard(userPassword, getCardID)
	},
}

var getCardID string //nolint:gochecknoglobals // cobra style guide

func init() {
	rootCmd.AddCommand(getCard)
	getCard.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	getCard.Flags().StringVarP(&getCardID, "card", "c", "", "Card id")

	if err := getCard.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := getCard.MarkFlagRequired("card"); err != nil {
		log.Fatal(err)
	}
}
