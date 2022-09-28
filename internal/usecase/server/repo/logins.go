package repo

import (
	"context"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/usecase/server/repo/models"
	"github.com/dimk00z/GophKeeper/internal/utils/errs"
	"github.com/google/uuid"
)

func (r *GophKeeperRepo) GetLogins(ctx context.Context, user entity.User) (logins []entity.Login, err error) {
	var loginsFromDB []models.SavedLogin

	err = r.db.WithContext(ctx).Find(&loginsFromDB, "user_id = ?", user.ID).Error
	if err != nil {
		return nil, err
	}

	if len(loginsFromDB) == 0 {
		return nil, nil
	}

	logins = make([]entity.Login, len(loginsFromDB))

	for index := range loginsFromDB {
		logins[index].ID = loginsFromDB[index].ID
		logins[index].Name = loginsFromDB[index].Name
		logins[index].Password = loginsFromDB[index].Password
		logins[index].URI = loginsFromDB[index].URI
		logins[index].Login = loginsFromDB[index].Login
	}

	return logins, nil
}

func (r *GophKeeperRepo) AddLogin(ctx context.Context, login *entity.Login, userID uuid.UUID) error {
	loginToDB := models.SavedLogin{
		ID:       uuid.New(),
		UserID:   userID,
		Name:     login.Name,
		Password: login.Password,
		URI:      login.URI,
		Login:    login.Login,
	}

	if err := r.db.WithContext(ctx).Create(&loginToDB).Error; err != nil {
		return err
	}

	login.ID = loginToDB.ID

	return nil
}

func (r *GophKeeperRepo) IsLoginOwner(ctx context.Context, loginID, userID uuid.UUID) bool {
	var loginFromDB models.SavedLogin

	r.db.WithContext(ctx).Where("id = ?", loginID).First(&loginFromDB)

	return loginFromDB.UserID == userID
}

func (r *GophKeeperRepo) DelLogin(ctx context.Context, loginID, userID uuid.UUID) error {
	if !r.IsLoginOwner(ctx, loginID, userID) {
		return errs.ErrWrongOwnerOrNotFound
	}

	return r.db.WithContext(ctx).Delete(&models.SavedLogin{}, loginID).Error
}

func (r *GophKeeperRepo) UpdateLogin(ctx context.Context, login *entity.Login, userID uuid.UUID) error {
	if !r.IsLoginOwner(ctx, login.ID, userID) {
		return errs.ErrWrongOwnerOrNotFound
	}

	loginToDB := models.SavedLogin{
		ID:       login.ID,
		Name:     login.Name,
		Password: login.Password,
		URI:      login.URI,
		Login:    login.Login,
		UserID:   userID,
	}

	return r.db.WithContext(ctx).Save(&loginToDB).Error
}
