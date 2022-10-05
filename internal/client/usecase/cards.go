package usecase

import (
	"log"

	"github.com/fatih/color"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/utils"
)

func (uc *GophKeeperClientUseCase) AddCard(userPassword string, card *entity.Card) {
	if !uc.verifyPassword(userPassword) {
		return
	}
	accessToken, err := uc.repo.GetSavedAccessToken()
	if err != nil || accessToken == "" {
		color.Red("User should be logged")

		return
	}
	uc.encryptCard(userPassword, card)

	if err = uc.clientAPI.AddCard(accessToken, card); err != nil {
		return
	}

	uc.repo.AddCard(card)

	color.Green("Card %q added, id: %v", card.Name, card.ID)
}

func (uc *GophKeeperClientUseCase) encryptCard(userPassword string, card *entity.Card) {
	card.Number = utils.Encrypt(userPassword, card.Number)
	if card.SecurityCode != "" {
		card.SecurityCode = utils.Encrypt(userPassword, card.SecurityCode)
	}
	if card.ExpirationMonth != "" {
		card.ExpirationMonth = utils.Encrypt(userPassword, card.ExpirationMonth)
	}
	if card.ExpirationYear != "" {
		card.ExpirationYear = utils.Encrypt(userPassword, card.ExpirationYear)
	}
	if card.CardHolderName != "" {
		card.CardHolderName = utils.Encrypt(userPassword, card.CardHolderName)
	}
}

func (uc *GophKeeperClientUseCase) decryptCard(userPassword string, card *entity.Card) {
	card.Number = utils.Decrypt(userPassword, card.Number)
	if card.SecurityCode != "" {
		card.SecurityCode = utils.Decrypt(userPassword, card.SecurityCode)
	}
	if card.ExpirationMonth != "" {
		card.ExpirationMonth = utils.Decrypt(userPassword, card.ExpirationMonth)
	}
	if card.ExpirationYear != "" {
		card.ExpirationYear = utils.Decrypt(userPassword, card.ExpirationYear)
	}
	if card.CardHolderName != "" {
		card.CardHolderName = utils.Decrypt(userPassword, card.CardHolderName)
	}
}

func (uc *GophKeeperClientUseCase) loadCards(accessToken string) {
	cards, err := uc.clientAPI.GetCards(accessToken)
	if err != nil {
		color.Red("Connection error: %v", err)

		return
	}

	if err = uc.repo.SaveCards(cards); err != nil {
		log.Println(err)

		return
	}
	color.Green("Loaded %v cards", len(cards))
}
