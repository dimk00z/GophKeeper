package usecase

import (
	"context"
	"mime/multipart"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/utils"
	"github.com/google/uuid"
)

func (uc *GophKeeperUseCase) GetBinaries(ctx context.Context, user entity.User) ([]entity.Binary, error) {
	return uc.repo.GetBinaries(ctx, user)
}

func (uc *GophKeeperUseCase) AddBinary(
	ctx context.Context,
	binary *entity.Binary,
	file *multipart.FileHeader,
	userID uuid.UUID,
) error {
	userDirectory := uc.cfg.FilesStorage.Location + "/" + userID.String()

	if err := uc.repo.AddBinary(ctx, binary, userID); err != nil {
		return err
	}

	if err := utils.SaveUploadedFile(file, binary.ID.String(), userDirectory); err != nil {
		uc.l.Debug("GophKeeperUseCase - AddBinary - SaveUploadedFile - %w", err)

		return err
	}

	return nil
}
