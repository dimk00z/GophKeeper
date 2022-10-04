package usecase

import (
	config "github.com/dimk00z/GophKeeper/config/client"
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
