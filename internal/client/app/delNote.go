package app

import (
	"log"

	"github.com/spf13/cobra"
)

var delNote = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "delnote",
	Short: "Delete user note by id",
	Long: `
This command remove note
Usage: delnote -i \"note_id\" 
Flags:
  -i, --id string Card id
  -p, --password string   User password value.`,
	Run: func(cmd *cobra.Command, args []string) {
		clientUseCase.DelNote(userPassword, delNoteID)
	},
}

var delNoteID string //nolint:gochecknoglobals // cobra style guide

func init() {
	rootCmd.AddCommand(delNote)
	delNote.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	delNote.Flags().StringVarP(&delNoteID, "id", "i", "", "Note id")

	if err := delNote.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := delNote.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
