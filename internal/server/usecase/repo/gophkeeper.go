package repo

import (
	"github.com/dimk00z/GophKeeper/internal/server/usecase/repo/models"
	"github.com/dimk00z/GophKeeper/pkg/logger"
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
		&models.Card{},
		&models.Login{},
		&models.Note{},
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

func (r *GophKeeperRepo) ShutDown() {
	db, err := r.db.DB()
	if err != nil {
		r.l.Error(err)
	}

	db.Close()
	r.l.Debug("db connection closed")
}
