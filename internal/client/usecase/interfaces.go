package usecase

import "github.com/dimk00z/GophKeeper/internal/entity"

type (
	// GophKeeperClient - use cases.
	GophKeeperClient interface {
		InitDB()
		Register(user *entity.User)
		Login(user *entity.User)
	}
	GophKeeperClientRepo interface {
		MigrateDB()
		AddUser(user *entity.User) error
		UpdateUserToken(user *entity.User, token *entity.JWT) error
		RemoveUsers()
		UserExistsByEmail(email string) bool
	}
)
