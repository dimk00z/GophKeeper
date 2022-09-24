package repo

import (
	"context"
	"fmt"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/usecase/repo/models"
	"github.com/dimk00z/GophKeeper/internal/utils/errs"
	"github.com/dimk00z/GophKeeper/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const _defaultEntityCap = 64

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
		l:  l}
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
	// TODO: add logic
	newUser := models.User{
		Email:    email,
		Password: hashedPassword,
	}
	result := r.db.Create(&newUser)

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

// GetHistory -.
func (r *GophKeeperRepo) GetHistory(ctx context.Context) ([]entity.GophKeeper, error) {
	// sql, _, err := r.Builder.
	// 	Select("source, destination, original, GophKeeper").
	// 	From("history").
	// 	ToSql()
	// if err != nil {
	// 	return nil, fmt.Errorf("GophKeeperRepo - GetHistory - r.Builder: %w", err)
	// }

	// rows, err := r.Pool.Query(ctx, sql)
	// if err != nil {
	// 	return nil, fmt.Errorf("GophKeeperRepo - GetHistory - r.Pool.Query: %w", err)
	// }
	// defer rows.Close()

	// entities := make([]entity.GophKeeper, 0, _defaultEntityCap)

	// for rows.Next() {
	// 	e := entity.GophKeeper{}

	// 	err = rows.Scan(&e.Source, &e.Destination, &e.Original, &e.GophKeeper)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("GophKeeperRepo - GetHistory - rows.Scan: %w", err)
	// 	}

	// 	entities = append(entities, e)
	// }

	return nil, nil
}

// Store -.
func (r *GophKeeperRepo) Store(ctx context.Context, t entity.GophKeeper) error {
	// sql, args, err := r.Builder.
	// 	Insert("history").
	// 	Columns("source, destination, original, GophKeeper").
	// 	Values(t.Source, t.Destination, t.Original, t.GophKeeper).
	// 	ToSql()
	// if err != nil {
	// 	return fmt.Errorf("GophKeeperRepo - Store - r.Builder: %w", err)
	// }

	// _, err = r.Pool.Exec(ctx, sql, args...)
	// if err != nil {
	// 	return fmt.Errorf("GophKeeperRepo - Store - r.Pool.Exec: %w", err)
	// }

	return nil
}
