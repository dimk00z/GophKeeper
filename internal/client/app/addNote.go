package app

import (
	"log"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/spf13/cobra"
)

var addNote = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "addnote",
	Short: "Add note",
	Long: `
This command add user note
Usage: addnote -p \"user_password\" 
Flags:
  -h, --help              help for addlogin
  -p, --password string   User password value.
  -n, --note string     User note`,
	Run: func(cmd *cobra.Command, args []string) {
		clientUseCase.AddNote(userPassword, &noteForAdditing)
	},
}

var noteForAdditing entity.SecretNote //nolint:gochecknoglobals // cobra style guide

func init() {
	rootCmd.AddCommand(addNote)
	addNote.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")

	addNote.Flags().StringVarP(&noteForAdditing.Name, "title", "t", "", "Login title")
	addNote.Flags().StringVarP(&noteForAdditing.Note, "note", "n", "", "User note")

	if err := addNote.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := addNote.MarkFlagRequired("title"); err != nil {
		log.Fatal(err)
	}
}
