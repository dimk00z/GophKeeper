package usecase

import (
	"errors"

	config "github.com/dimk00z/GophKeeper/config/client"
	"github.com/fatih/color"
)

type GophKeeperClientUseCase struct {
	repo      GophKeeperClientRepo
	clientAPI GophKeeperClientAPI
	cfg       *config.Config
}

func New(repo GophKeeperClientRepo, clientAPI GophKeeperClientAPI, cfg *config.Config) *GophKeeperClientUseCase {
	return &GophKeeperClientUseCase{
		repo:      repo,
		cfg:       cfg,
		clientAPI: clientAPI,
	}
}

func (uc *GophKeeperClientUseCase) InitDB() {
	uc.repo.MigrateDB()
}

var (
	errPasswordCheck = errors.New("wrong password")
	errToken         = errors.New("user token erroe")
)

func (uc *GophKeeperClientUseCase) authorisationCheck(userPassword string) (string, error) {
	if !uc.verifyPassword(userPassword) {
		return "", errPasswordCheck
	}
	accessToken, err := uc.repo.GetSavedAccessToken()
	if err != nil || accessToken == "" {
		color.Red("User should be logged")

		return "", errToken
	}

	return accessToken, nil
}
