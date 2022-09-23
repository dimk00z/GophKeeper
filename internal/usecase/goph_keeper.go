package usecase

import (
	"context"
	"fmt"

	"github.com/dimk00z/GophKeeper/internal/entity"
)

// GophKeeperUseCase -.
type GophKeeperUseCase struct {
	repo   GophKeeperRepo
	webAPI GophKeeperWebAPI
}

// New -.
func New(r GophKeeperRepo, w GophKeeperWebAPI) *GophKeeperUseCase {
	return &GophKeeperUseCase{
		repo:   r,
		webAPI: w,
	}
}

// History - getting translate history from store.
func (uc *GophKeeperUseCase) History(ctx context.Context) ([]entity.GophKeeper, error) {
	GophKeepers, err := uc.repo.GetHistory(ctx)
	if err != nil {
		return nil, fmt.Errorf("GophKeeperUseCase - History - s.repo.GetHistory: %w", err)
	}

	return GophKeepers, nil
}

// Translate -.
func (uc *GophKeeperUseCase) Translate(ctx context.Context, t entity.GophKeeper) (entity.GophKeeper, error) {
	GophKeeper, err := uc.webAPI.Translate(t)
	if err != nil {
		return entity.GophKeeper{}, fmt.Errorf("GophKeeperUseCase - Translate - s.webAPI.Translate: %w", err)
	}

	err = uc.repo.Store(context.Background(), GophKeeper)
	if err != nil {
		return entity.GophKeeper{}, fmt.Errorf("GophKeeperUseCase - Translate - s.repo.Store: %w", err)
	}

	return GophKeeper, nil
}

func (uc *GophKeeperUseCase) HealthCheck() error {
	return uc.repo.DBHealthCheck()
}

func (uc *GophKeeperUseCase) SignUpUser(ctx context.Context, email, password string) (user entity.User, err error) {
	// TODO: add logic
	return
}
