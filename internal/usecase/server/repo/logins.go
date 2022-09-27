package repo

import (
	"context"

	"github.com/dimk00z/GophKeeper/internal/entity"
)

func (r *GophKeeperRepo) GetLogins(ctx context.Context, user entity.User) (logins []entity.Login, err error) {
	return
}
