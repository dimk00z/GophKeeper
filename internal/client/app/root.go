package app

import (
	"fmt"
	"os"

	config "github.com/dimk00z/GophKeeper/config/client"
	"github.com/dimk00z/GophKeeper/internal/client/app/build"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   config.LoadConfig().App.Name,
	Short: "App for storing private data",
	Long:  `User can save cards, note and logins`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		build.PrintBulidInfo()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
