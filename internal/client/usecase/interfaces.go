package usecase

import (
	"github.com/dimk00z/GophKeeper/internal/client/usecase/viewsets"
	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/google/uuid"
)

type (
	// GophKeeperClient - use cases.
	GophKeeperClient interface {
		InitDB()

		Register(user *entity.User)
		Login(user *entity.User)
		Logout()
		Sync(userPassword string)

		ShowVault(userPassword, showVaultOption string)

		AddCard(userPassword string, card *entity.Card)
		DetailCardByID(userPassword string, cardID uuid.UUID)
	}
	GophKeeperClientRepo interface {
		MigrateDB()
		AddUser(user *entity.User) error
		UpdateUserToken(user *entity.User, token *entity.JWT) error
		DropUserToken() error
		GetSavedAccessToken() (string, error)
		RemoveUsers()
		UserExistsByEmail(email string) bool
		GetUserPasswordHash() string

		AddCard(*entity.Card)
		SaveCards([]entity.Card) error
		LoadCards() []viewsets.CardForList
		GetCardByID(cardID uuid.UUID) entity.Card

		AddLogin(*entity.Login)
		SaveLogins([]entity.Login) error
		LoadLogins() []viewsets.LoginForList
		GetLoginByID(loginID uuid.UUID) entity.Login

		LoadNotes() []viewsets.NoteForList
	}
	GophKeeperClientAPI interface {
		Login(user *entity.User) (entity.JWT, error)
		Register(user *entity.User) error
		GetCards(accessToken string) ([]entity.Card, error)
		AddCard(accessToken string, card *entity.Card) error
		GetLogins(accessToken string) ([]entity.Login, error)
		GetNotes(accessToken string) ([]entity.SecretNote, error)
	}
)
