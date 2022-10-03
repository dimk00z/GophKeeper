package app

import (
	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/spf13/cobra"
)

var registerUserCmd = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "register",
	Short: "Register user to the service",
	Long: `
This command register new user.
Usage: gophkeeperclient register user_login user_password`,
	Args: cobra.MinimumNArgs(RequiredUserArgs),
	Run: func(cmd *cobra.Command, args []string) {
		account := entity.User{
			Email:    args[0],
			Password: args[1],
		}

		clientUseCase.Register(&account)
	},
}

func init() {
	rootCmd.AddCommand(registerUserCmd)
}
