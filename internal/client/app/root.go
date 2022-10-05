package app

import (
	"fmt"
	"log"
	"os"

	config "github.com/dimk00z/GophKeeper/config/client"
	"github.com/dimk00z/GophKeeper/internal/client/app/build"
	"github.com/dimk00z/GophKeeper/internal/client/usecase"
	api "github.com/dimk00z/GophKeeper/internal/client/usecase/client_api"
	"github.com/dimk00z/GophKeeper/internal/client/usecase/repo"
	"github.com/spf13/cobra"
)

var (
	cfg           *config.Config           //nolint:gochecknoglobals // cobra style guide
	clientUseCase usecase.GophKeeperClient //nolint:gochecknoglobals // cobra style guide

	rootCmd = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
		Use:   config.LoadConfig().App.Name,
		Short: "App for storing private data",
		Long:  `User can save cards, note and logins`,
		Run: func(cmd *cobra.Command, args []string) {
			build.PrintBulidInfo()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initApp)
}

func initApp() {
	cfg = config.LoadConfig()

	log.Println(cfg)

	clientUseCase = usecase.New(repo.New(cfg.SQLite.DSN), api.New(cfg.Server.URL), cfg)
}
