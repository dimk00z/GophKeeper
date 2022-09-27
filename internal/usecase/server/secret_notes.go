package usecase

import (
	"context"

	"github.com/dimk00z/GophKeeper/internal/entity"
)

func (uc *GophKeeperUseCase) GetSecretNotes(ctx context.Context, user entity.User) ([]entity.SecretNote, error) {
	return uc.repo.GetSecretNotes(ctx, user)
}
