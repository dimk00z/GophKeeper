package usecase

import (
	config "github.com/dimk00z/GophKeeper/config/client"
)

type GophKeeperClientUseCase struct {
	repo GophKeeperClientRepo
	cfg  *config.Config
}

func New(repo GophKeeperClientRepo, cfg *config.Config) *GophKeeperClientUseCase {
	return &GophKeeperClientUseCase{
		repo: repo,
		cfg:  cfg,
	}
}

func (uc *GophKeeperClientUseCase) InitDB() {
	uc.repo.MigrateDB()
}
