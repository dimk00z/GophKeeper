package app

import (
	"github.com/dimk00z/GophKeeper/internal/client/usecase"
	"github.com/spf13/cobra"
)

var RegisterInitLocalStorage = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "init",
	Short: "Init local storage",
	Long: `
This command register init sqlite db for storaging private data.
Usage: gophkeeperclient init`,
	Run: func(cmd *cobra.Command, args []string) {
		usecase.GetClientUseCase().InitDB()
	},
}
