package app

import (
	"github.com/spf13/cobra"
)

var registerInitLocalStorage = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "init",
	Short: "Init local storage",
	Long: `
This command register init sqlite db for storaging private data.
Usage: gophkeeperclient init`,
	Run: func(cmd *cobra.Command, args []string) {
		clientUseCase.InitDB()
	},
}

func init() {
	rootCmd.AddCommand(registerInitLocalStorage)
}
