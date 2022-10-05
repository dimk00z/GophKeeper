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
		ShowCard(userPassword, cardID string)
		DelCard(userPassword, cardID string)

		AddLogin(userPassword string, login *entity.Login)
		ShowLogin(userPassword, loginID string)
		DelLogin(userPassword, loginID string)

		AddNote(userPassword string, note *entity.SecretNote)
		ShowNote(userPassword, noteID string)
		DelNote(userPassword, noteID string)
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
		GetCardByID(cardID uuid.UUID) (entity.Card, error)
		DelCard(cardID uuid.UUID) error

		AddLogin(*entity.Login)
		SaveLogins([]entity.Login) error
		LoadLogins() []viewsets.LoginForList
		GetLoginByID(loginID uuid.UUID) (entity.Login, error)
		DelLogin(loginID uuid.UUID) error

		LoadNotes() []viewsets.NoteForList
		SaveNotes([]entity.SecretNote) error
		AddNote(*entity.SecretNote)
		GetNoteByID(notedID uuid.UUID) (entity.SecretNote, error)
		DelNote(noteID uuid.UUID) error
	}
	GophKeeperClientAPI interface {
		Login(user *entity.User) (entity.JWT, error)
		Register(user *entity.User) error

		GetCards(accessToken string) ([]entity.Card, error)
		AddCard(accessToken string, card *entity.Card) error
		DelCard(accessToken, cardID string) error

		GetLogins(accessToken string) ([]entity.Login, error)
		AddLogin(accessToken string, login *entity.Login) error
		DelLogin(accessToken, loginID string) error

		GetNotes(accessToken string) ([]entity.SecretNote, error)
		AddNote(accessToken string, note *entity.SecretNote) error
		DelNote(accessToken, noteID string) error
	}
)
