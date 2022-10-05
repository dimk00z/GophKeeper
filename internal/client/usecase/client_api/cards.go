package clientapi

import (
	"github.com/dimk00z/GophKeeper/internal/entity"
)

const cardsEndpoint = "api/v1/user/cards"

func (api *GophKeeperClientAPI) GetCards(accessToken string) (cards []entity.Card, err error) {
	if err := api.getEntities(&cards, accessToken, cardsEndpoint); err != nil {
		return nil, err
	}

	return cards, nil
}

func (api *GophKeeperClientAPI) AddCard(accessToken string, card *entity.Card) error {
	return api.addEntity(card, accessToken, cardsEndpoint)
}
