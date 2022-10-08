package app

import (
	"log"

	"github.com/dimk00z/GophKeeper/internal/entity"
	utils "github.com/dimk00z/GophKeeper/internal/utils/client"
	"github.com/spf13/cobra"
)

var addCard = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "addcard",
	Short: "Add card",
	Long: `
This command add card
Usage: addcard -p \"user_password\" 
Flags:
  -b, --brand string      Card brand
  -c, --code string       Card code
  -h, --help              help for addcard
  -m, --month string      Card expiration month
  -n, --number string     Card namber
  -o, --owner string      Card holder name
  -p, --password string   User password value.
  -t, --title string      Card title
  -y, --year string       Card expiration year
  -meta 				  Add meta data for entiry
  example: -meta'[{"name":"some_meta","value":"some_meta_value"},{"name":"some_meta2","value":"some_meta_value2"}]'
  `,
	Run: func(cmd *cobra.Command, args []string) {
		clientUseCase.AddCard(userPassword, &cardForAdditing)
	},
}

var cardForAdditing entity.Card //nolint:gochecknoglobals // cobra style guide

func init() {
	rootCmd.AddCommand(addCard)
	addCard.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	addCard.Flags().StringVarP(&cardForAdditing.Name, "title", "t", "", "Card title")
	addCard.Flags().StringVarP(&cardForAdditing.Number, "number", "n", "", "Card namber")
	addCard.Flags().StringVarP(&cardForAdditing.CardHolderName, "owner", "o", "", "Card holder name")
	addCard.Flags().StringVarP(&cardForAdditing.Brand, "brand", "b", "", "Card brand")
	addCard.Flags().StringVarP(&cardForAdditing.SecurityCode, "code", "c", "", "Card code")
	addCard.Flags().StringVarP(&cardForAdditing.ExpirationMonth, "month", "m", "", "Card expiration month")
	addCard.Flags().StringVarP(&cardForAdditing.ExpirationYear, "year", "y", "", "Card expiration year")
	addCard.Flags().Var(&utils.JSONFlag{Target: &cardForAdditing.Meta}, "meta", `Add meta fields for entity`)

	if err := addCard.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := addCard.MarkFlagRequired("title"); err != nil {
		log.Fatal(err)
	}
	if err := addCard.MarkFlagRequired("number"); err != nil {
		log.Fatal(err)
	}
	if err := addCard.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
}
