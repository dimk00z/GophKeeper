package app

import (
	"log"

	"github.com/spf13/cobra"
)

var syncUserData = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "sync",
	Short: "Sync user`s data",
	Long: `
This command update users private data from server
Usage: gophkeeperclient sync -p \"user_password\"`,
	Run: func(cmd *cobra.Command, args []string) {
		clientUseCase.Sync(userPassword)
	},
}

var userPassword string //nolint:gochecknoglobals // cobra style guide

func init() {
	rootCmd.AddCommand(syncUserData)
	syncUserData.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	if err := syncUserData.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
}
