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
	}

	// GophKeeperRepo - db logic.
	GophKeeperRepo interface {
		DBHealthCheck() error
		AddUser(ctx context.Context, email, hashedPassword string) (entity.User, error)
		GetUser(ctx context.Context, email, hashedPassword string) (entity.User, error)
	}

	// GophKeeperWebAPI - business logic.
	GophKeeperWebAPI interface {
		Translate(entity.GophKeeper) (entity.GophKeeper, error)
	}
)
