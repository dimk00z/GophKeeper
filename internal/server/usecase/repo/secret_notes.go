package repo

import (
	"context"

	"github.com/dimk00z/GophKeeper/internal/entity"
)

func (r *GophKeeperRepo) GetSecretNotes(ctx context.Context, user entity.User) (notes []entity.SecretNote, err error) {
	return
}
