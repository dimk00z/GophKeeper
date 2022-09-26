package repo

import (
	"context"
	"fmt"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/usecase/repo/models"
	"github.com/dimk00z/GophKeeper/internal/utils"
	"github.com/dimk00z/GophKeeper/internal/utils/errs"
	"github.com/dimk00z/GophKeeper/pkg/logger"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GophKeeper -.
type GophKeeperRepo struct {
	db *gorm.DB
	l  *logger.Logger
}

// New -.
func New(dsn string, l *logger.Logger) *GophKeeperRepo {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Fatal("Repo - new - %v", err)
	}

	return &GophKeeperRepo{
		db: db,
		l:  l,
	}
}

func (r *GophKeeperRepo) Migrate() {
	tables := []interface{}{
		&models.User{},
		&models.CreditCard{},
		&models.SavedLogin{},
	}

	if err := r.db.AutoMigrate(tables...); err != nil {
		r.l.Fatal("GophKeeperRepo - Migrate - %v", err)
	}

	r.l.Debug("GophKeeperRepo - Migrate - success")
}

func (r *GophKeeperRepo) DBHealthCheck() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}

func (r *GophKeeperRepo) AddUser(ctx context.Context, email, hashedPassword string) (user entity.User, err error) {
	newUser := models.User{
		Email:    email,
		Password: hashedPassword,
	}
	result := r.db.WithContext(ctx).Create(&newUser)

	if result.Error == nil {
		user.ID = newUser.ID
		user.Email = newUser.Email

		return
	}

	switch errs.ParsePostgresErr(result.Error).Code {
	case "23505":
		r.l.Debug("AddUser - %w", result.Error)

		err = errs.ErrEmailAlreadyExists

		return
	default:
		err = fmt.Errorf("AddUser - %w", result.Error)

		return
	}
}

func (r *GophKeeperRepo) GetUserByEmail(ctx context.Context, email, hashedPassword string) (user entity.User, err error) {
	var userFromDB models.User

	r.db.WithContext(ctx).Where("email = ?", email).First(&userFromDB)

	if userFromDB.ID == uuid.Nil {
		err = errs.ErrWrongCredentials

		return
	}

	if err = utils.VerifyPassword(userFromDB.Password, hashedPassword); err != nil {
		err = errs.ErrWrongCredentials

		return
	}

	user.ID = userFromDB.ID
	user.Email = userFromDB.Email

	return
}

func (r *GophKeeperRepo) GetUserByID(ctx context.Context, id string) (user entity.User, err error) {
	var userFromDB models.User

	r.db.WithContext(ctx).Where("id = ?", id).First(&userFromDB)

	if userFromDB.ID == uuid.Nil {
		err = errs.ErrWrongCredentials

		return
	}

	user.ID = userFromDB.ID
	user.Email = userFromDB.Email

	return
}

func (r *GophKeeperRepo) ShutDown() {
	db, err := r.db.DB()
	if err != nil {
		r.l.Error(err)
	}

	db.Close()
	r.l.Debug("db connection closed")
}
