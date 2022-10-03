package repo

import (
	"github.com/dimk00z/GophKeeper/internal/client/usecase/repo/models"
	"github.com/fatih/color"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GophKeeperRepo struct {
	db *gorm.DB
}

func New(dbFileName string) *GophKeeperRepo {
	db, err := gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	if err != nil {
		color.Red("Load error %s", err.Error())
	}

	return &GophKeeperRepo{
		db: db,
	}
}

func (r *GophKeeperRepo) MigrateDB() {
	tables := []interface{}{
		&models.User{},
		&models.Card{},
		&models.Login{},
		&models.Note{},
	}
	var err error
	for _, table := range tables {

		if err = r.db.Migrator().DropTable(table); err != nil {
			color.Red("Init error %s", err.Error())
		}
		if err = r.db.Migrator().CreateTable(table); err != nil {
			color.Red("Init error %s", err.Error())
		}
	}

	color.Green("Initialization status: success")
	color.Green("You can use gophkeer")
}
