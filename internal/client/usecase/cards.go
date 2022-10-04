package usecase

import (
	"log"

	"github.com/fatih/color"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/utils"
	"github.com/google/uuid"
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

	color.Green("Card %q added", card.Name)
}

func (uc *GophKeeperClientUseCase) encryptCard(userPassword string, card *entity.Card) {
	card.Number = utils.Encrypt(userPassword, card.Number)
	card.SecurityCode = utils.Encrypt(userPassword, card.SecurityCode)
	card.ExpirationMonth = utils.Encrypt(userPassword, card.ExpirationMonth)
	card.ExpirationYear = utils.Encrypt(userPassword, card.ExpirationYear)
	card.CardHolderName = utils.Encrypt(userPassword, card.CardHolderName)
}

func (uc *GophKeeperClientUseCase) decryptCard(userPassword string, card *entity.Card) {
	card.Number = utils.Decrypt(userPassword, card.Number)
	card.SecurityCode = utils.Decrypt(userPassword, card.SecurityCode)
	card.ExpirationMonth = utils.Decrypt(userPassword, card.ExpirationMonth)
	card.ExpirationYear = utils.Decrypt(userPassword, card.ExpirationYear)
	card.CardHolderName = utils.Decrypt(userPassword, card.CardHolderName)
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

func (uc *GophKeeperClientUseCase) ListCards() {
	// TODO add logic
}

func (uc *GophKeeperClientUseCase) DetailCardByID(userPassword string, cardID uuid.UUID) {
	// TODO add logic
}

func (uc *GophKeeperClientUseCase) DetailCardByName(userPassword, cardName string) {
	// TODO add logic
}
