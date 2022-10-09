package repo

import (
	"context"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/server/usecase/repo/models"
	"github.com/google/uuid"
)

func (r *GophKeeperRepo) GetBinaries(ctx context.Context, user entity.User) ([]entity.Binary, error) {
	var binariesFromDB []models.Binary

	if err := r.db.WithContext(ctx).Find(&binariesFromDB, "user_id = ?", user.ID).Error; err != nil {
		return nil, err
	}

	if len(binariesFromDB) == 0 {
		return nil, nil
	}

	binaries := make([]entity.Binary, len(binariesFromDB))

	for index := range binariesFromDB {
		binaries[index].ID = binariesFromDB[index].ID
		binaries[index].Name = binariesFromDB[index].Name
		binaries[index].FileName = binariesFromDB[index].FileName
	}

	return binaries, nil
}

func (r *GophKeeperRepo) AddBinary(ctx context.Context, binary *entity.Binary, userID uuid.UUID) error {
	newBinaryToDB := models.Binary{
		Name:     binary.Name,
		FileName: binary.FileName,
		UserID:   userID,
	}

	if err := r.db.WithContext(ctx).Create(&newBinaryToDB).Error; err != nil {
		r.l.Debug("GophKeeperRepo - AddBinary - Create - %w", err)

		return err
	}
	binary.ID = newBinaryToDB.ID

	return nil
}
