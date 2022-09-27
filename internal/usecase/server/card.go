package usecase

import (
	"context"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/google/uuid"
)

func (uc *GophKeeperUseCase) GetCards(ctx context.Context, user entity.User) ([]entity.Card, error) {
	return uc.repo.GetCards(ctx, user)
}

func (uc *GophKeeperUseCase) AddCard(ctx context.Context, card *entity.Card, userID uuid.UUID) error {
	return uc.repo.AddCard(ctx, card, userID)
}
