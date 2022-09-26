// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/dimk00z/GophKeeper/internal/entity"
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
	}

	// GophKeeperRepo - db logic.
	GophKeeperRepo interface {
		DBHealthCheck() error
		AddUser(ctx context.Context, email, hashedPassword string) (entity.User, error)
		GetUserByEmail(ctx context.Context, email, hashedPassword string) (entity.User, error)
		GetUserByID(ctx context.Context, id string) (entity.User, error)
	}

	// GophKeeperWebAPI - business logic.
	GophKeeperWebAPI interface {
		// Translate(entity.GophKeeper) (entity.GophKeeper, error)
	}
)
