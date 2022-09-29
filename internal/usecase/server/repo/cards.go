package repo

import (
	"context"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/usecase/server/repo/models"
	"github.com/dimk00z/GophKeeper/internal/utils/errs"
	"github.com/google/uuid"
)

func (r *GophKeeperRepo) GetCards(ctx context.Context, user entity.User) ([]entity.Card, error) {
	var cardsFromDB []models.Card

	err := r.db.WithContext(ctx).Find(&cardsFromDB, "user_id = ?", user.ID).Error
	if err != nil {
		return nil, err
	}

	if len(cardsFromDB) == 0 {
		return nil, nil
	}

	cards := make([]entity.Card, len(cardsFromDB))

	for index := range cardsFromDB {
		cards[index].ID = cardsFromDB[index].ID
		cards[index].Brand = cardsFromDB[index].Brand
		cards[index].CardHolderName = cardsFromDB[index].CardHolderName
		cards[index].ExpirationMonth = cardsFromDB[index].ExpirationMonth
		cards[index].ExpirationYear = cardsFromDB[index].ExpirationYear
		cards[index].Name = cardsFromDB[index].Name
		cards[index].Number = cardsFromDB[index].Number
		cards[index].SecurityCode = cardsFromDB[index].SecurityCode
	}

	return cards, nil
}

func (r *GophKeeperRepo) AddCard(ctx context.Context, card *entity.Card, userID uuid.UUID) error {
	cardToDB := models.Card{
		ID:              uuid.New(),
		UserID:          userID,
		Name:            card.Name,
		Brand:           card.Brand,
		CardHolderName:  card.CardHolderName,
		Number:          card.Number,
		ExpirationMonth: card.ExpirationMonth,
		ExpirationYear:  card.ExpirationYear,
		SecurityCode:    card.SecurityCode,
	}

	if err := r.db.WithContext(ctx).Create(&cardToDB).Error; err != nil {
		return err
	}

	card.ID = cardToDB.ID

	return nil
}

func (r *GophKeeperRepo) IsCardOwner(ctx context.Context, cardUUID, userID uuid.UUID) bool {
	var cardFromDB models.Card

	r.db.WithContext(ctx).Where("id = ?", cardUUID).First(&cardFromDB)

	return cardFromDB.UserID == userID
}

func (r *GophKeeperRepo) DelCard(ctx context.Context, cardUUID, userID uuid.UUID) error {
	if !r.IsCardOwner(ctx, cardUUID, userID) {
		return errs.ErrWrongOwnerOrNotFound
	}

	return r.db.WithContext(ctx).Delete(&models.Card{}, cardUUID).Error
}

func (r *GophKeeperRepo) UpdateCard(ctx context.Context, card *entity.Card, userID uuid.UUID) error {
	if !r.IsCardOwner(ctx, card.ID, userID) {
		return errs.ErrWrongOwnerOrNotFound
	}

	cardToDB := models.Card{
		ID:              card.ID,
		UserID:          userID,
		Name:            card.Name,
		Brand:           card.Brand,
		CardHolderName:  card.CardHolderName,
		Number:          card.Number,
		ExpirationMonth: card.ExpirationMonth,
		ExpirationYear:  card.ExpirationYear,
		SecurityCode:    card.SecurityCode,
	}

	return r.db.WithContext(ctx).Save(&cardToDB).Error
}
