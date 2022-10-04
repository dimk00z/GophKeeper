package repo

import (
	"github.com/dimk00z/GophKeeper/internal/client/usecase/repo/models"
	"github.com/dimk00z/GophKeeper/internal/entity"
)

func (r *GophKeeperRepo) AddCard(card *entity.Card) {
	var user models.User
	r.db.First(&user)
	cardForSaving := models.Card{
		ID:              card.ID,
		Brand:           card.Brand,
		Name:            card.Name,
		Number:          card.Number,
		SecurityCode:    card.SecurityCode,
		CardHolderName:  card.CardHolderName,
		ExpirationMonth: card.ExpirationMonth,
		ExpirationYear:  card.ExpirationYear,
		UserID:          user.ID,
	}
	r.db.Save(&cardForSaving)
}

func (r *GophKeeperRepo) SaveCards(cards []entity.Card) error {
	var user models.User
	r.db.First(&user)
	cardsForDB := make([]models.Card, len(cards))
	for index := range cards {
		cardsForDB[index].ID = cards[index].ID
		cardsForDB[index].Brand = cards[index].Brand
		cardsForDB[index].CardHolderName = cards[index].CardHolderName
		cardsForDB[index].ExpirationMonth = cards[index].ExpirationMonth
		cardsForDB[index].ExpirationYear = cards[index].ExpirationYear
		cardsForDB[index].Name = cards[index].Name
		cardsForDB[index].Number = cards[index].Number
		cardsForDB[index].SecurityCode = cards[index].SecurityCode
		cardsForDB[index].UserID = user.ID
	}

	return r.db.Save(cardsForDB).Error
}
