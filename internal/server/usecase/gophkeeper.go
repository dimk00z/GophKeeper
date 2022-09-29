package usecase

import (
	config "github.com/dimk00z/GophKeeper/config/server"
	c "github.com/dimk00z/GophKeeper/pkg/cache"
)

const minutesPerHour = 60

// GophKeeperUseCase -.
type GophKeeperUseCase struct {
	repo  GophKeeperRepo
	cfg   *config.Config
	cache c.Cacher
}

func New(r GophKeeperRepo, cfg *config.Config, cache c.Cacher) *GophKeeperUseCase {
	return &GophKeeperUseCase{
		repo:  r,
		cfg:   cfg,
		cache: cache,
	}
}

func (uc *GophKeeperUseCase) HealthCheck() error {
	return uc.repo.DBHealthCheck()
}

func (uc *GophKeeperUseCase) GetDomainName() string {
	return uc.cfg.Secutiry.Domain
}
