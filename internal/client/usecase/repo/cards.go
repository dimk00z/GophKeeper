package repo

import (
	"errors"

	"github.com/dimk00z/GophKeeper/internal/client/usecase/repo/models"
	"github.com/dimk00z/GophKeeper/internal/client/usecase/viewsets"
	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/google/uuid"
)

var errCardNotFound = errors.New("card not found")

func (r *GophKeeperRepo) AddCard(card *entity.Card) {
	cardForSaving := models.Card{
		ID:              card.ID,
		Brand:           card.Brand,
		Name:            card.Name,
		Number:          card.Number,
		SecurityCode:    card.SecurityCode,
		CardHolderName:  card.CardHolderName,
		ExpirationMonth: card.ExpirationMonth,
		ExpirationYear:  card.ExpirationYear,
		UserID:          r.getUserID(),
	}
	r.db.Save(&cardForSaving)
}

func (r *GophKeeperRepo) SaveCards(cards []entity.Card) error {
	if len(cards) == 0 {
		return nil
	}
	userID := r.getUserID()
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
		cardsForDB[index].UserID = userID
	}

	return r.db.Save(cardsForDB).Error
}

func (r *GophKeeperRepo) LoadCards() []viewsets.CardForList {
	userID := r.getUserID()
	var cards []models.Card
	r.db.Where("user_id", userID).Find(&cards)
	if len(cards) == 0 {
		return nil
	}

	cardsViewSet := make([]viewsets.CardForList, len(cards))

	for index := range cards {
		cardsViewSet[index].ID = cards[index].ID
		cardsViewSet[index].Name = cards[index].Name
		cardsViewSet[index].Brand = cards[index].Brand
	}

	return cardsViewSet
}

func (r *GophKeeperRepo) GetCardByID(cardID uuid.UUID) (card entity.Card, err error) {
	var cardFromDB models.Card
	if err = r.db.Find(&cardFromDB, cardID).Error; cardFromDB.ID == uuid.Nil || err != nil {
		return card, errCardNotFound
	}

	card.ID = cardFromDB.ID
	card.Brand = cardFromDB.Brand
	card.Number = cardFromDB.Number
	card.Name = cardFromDB.Name
	card.CardHolderName = cardFromDB.CardHolderName
	card.ExpirationMonth = cardFromDB.ExpirationMonth
	card.ExpirationYear = cardFromDB.ExpirationYear

	return
}
