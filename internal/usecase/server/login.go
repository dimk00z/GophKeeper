package usecase

import (
	"context"

	"github.com/dimk00z/GophKeeper/internal/entity"
)

func (uc *GophKeeperUseCase) GetLogins(ctx context.Context, user entity.User) ([]entity.Login, error) {
	return uc.repo.GetLogins(ctx, user)
}
