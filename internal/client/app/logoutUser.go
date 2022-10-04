package app

import "github.com/spf13/cobra"

var logoutUser = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "logout",
	Short: "Logout user",
	Long: `
This command drops users tokens`,
	Run: func(cmd *cobra.Command, args []string) {
		clientUseCase.Logout()
	},
}

func init() {
	rootCmd.AddCommand(logoutUser)
}
