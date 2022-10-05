package app

import (
	"log"

	"github.com/spf13/cobra"
)

var getNote = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "getnote",
	Short: "Show user note by id",
	Long: `
This command show user note
Usage: getnote -i \"note_id\" 
Flags:
  -i, --id string Note id
  -p, --password string   User password value.`,
	Run: func(cmd *cobra.Command, args []string) {
		clientUseCase.ShowNote(userPassword, getNoteID)
	},
}

var getNoteID string //nolint:gochecknoglobals // cobra style guide

func init() {
	rootCmd.AddCommand(getNote)
	getNote.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	getNote.Flags().StringVarP(&getNoteID, "id", "i", "", "Note id")

	if err := getNote.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := getNote.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
