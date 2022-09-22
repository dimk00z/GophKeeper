// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/dimk00z/GophKeeper/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// GophKeeper -.
	GophKeeper interface {
		Translate(context.Context, entity.GophKeeper) (entity.GophKeeper, error)
		History(context.Context) ([]entity.GophKeeper, error)
	}

	// GophKeeperRepo -.
	GophKeeperRepo interface {
		Store(context.Context, entity.GophKeeper) error
		GetHistory(context.Context) ([]entity.GophKeeper, error)
	}

	// GophKeeperWebAPI -.
	GophKeeperWebAPI interface {
		Translate(entity.GophKeeper) (entity.GophKeeper, error)
	}
)
