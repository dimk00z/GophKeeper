package repo

import (
	"context"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/usecase/server/repo/models"
	"github.com/google/uuid"
)

func (r *GophKeeperRepo) GetCards(ctx context.Context, user entity.User) ([]entity.Card, error) {
	var cardsFromDB []models.CreditCard

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
	cardToDB := models.CreditCard{
		ID:              uuid.New(),
		UserID:          userID,
		Name:            card.Name,
		Brand:           card.Brand,
		CardHolderName:  card.CardHolderName,
		Number:          card.Name,
		ExpirationMonth: card.ExpirationMonth,
		ExpirationYear:  card.ExpirationYear,
		SecurityCode:    card.SecurityCode,
	}

	err := r.db.WithContext(ctx).Create(&cardToDB).Error
	if err != nil {
		card.ID = cardToDB.ID
	}

	return err
}
