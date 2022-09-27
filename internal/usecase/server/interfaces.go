// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/google/uuid"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// GophKeeper - use cases.
	GophKeeper interface {
		HealthCheck() error
		SignUpUser(ctx context.Context, email, password string) (entity.User, error)
		SignInUser(ctx context.Context, email, password string) (entity.JWT, error)
		RefreshAccessToken(ctx context.Context, refreshToken string) (entity.JWT, error)
		GetDomainName() string
		CheckAccessToken(ctx context.Context, accessToken string) (entity.User, error)

		GetLogins(ctx context.Context, user entity.User) ([]entity.Login, error)

		GetCards(ctx context.Context, user entity.User) ([]entity.Card, error)
		AddCard(ctx context.Context, card *entity.Card, userID uuid.UUID) error
		GetSecretNotes(ctx context.Context, user entity.User) ([]entity.SecretNote, error)
	}

	// GophKeeperRepo - db logic.
	GophKeeperRepo interface {
		DBHealthCheck() error
		AddUser(ctx context.Context, email, hashedPassword string) (entity.User, error)
		GetUserByEmail(ctx context.Context, email, hashedPassword string) (entity.User, error)
		GetUserByID(ctx context.Context, id string) (entity.User, error)

		GetLogins(ctx context.Context, user entity.User) ([]entity.Login, error)

		GetCards(ctx context.Context, user entity.User) ([]entity.Card, error)
		AddCard(ctx context.Context, card *entity.Card, userID uuid.UUID) error

		GetSecretNotes(ctx context.Context, user entity.User) ([]entity.SecretNote, error)
	}
)
