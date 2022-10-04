package usecase

import "github.com/dimk00z/GophKeeper/internal/entity"

type (
	// GophKeeperClient - use cases.
	GophKeeperClient interface {
		InitDB()
		Register(user *entity.User)
		Login(user *entity.User)
		Logout()
		Sync(userPassword string)
	}
	GophKeeperClientRepo interface {
		MigrateDB()
		AddUser(user *entity.User) error
		UpdateUserToken(user *entity.User, token *entity.JWT) error
		DropUserToken() error
		RemoveUsers()
		UserExistsByEmail(email string) bool
		GetUserPasswordHash() string
	}
	GophKeeperClientAPI interface {
		Login(user *entity.User) (entity.JWT, error)
		Register(user *entity.User) error
	}
)
