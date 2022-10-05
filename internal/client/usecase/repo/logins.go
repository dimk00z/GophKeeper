package repo

import (
	"errors"

	"github.com/dimk00z/GophKeeper/internal/client/usecase/repo/models"
	"github.com/dimk00z/GophKeeper/internal/client/usecase/viewsets"
	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/google/uuid"
)

var errLoginNotFound = errors.New("login not found")

func (r *GophKeeperRepo) AddLogin(login *entity.Login) {
	loginForSaving := models.Login{
		ID:       login.ID,
		Name:     login.Name,
		URI:      login.URI,
		Login:    login.Login,
		Password: login.Password,
		UserID:   r.getUserID(),
	}
	r.db.Save(&loginForSaving)
}

func (r *GophKeeperRepo) SaveLogins(logins []entity.Login) error {
	if len(logins) == 0 {
		return nil
	}
	userID := r.getUserID()
	loginsForDB := make([]models.Login, len(logins))
	for index := range logins {
		loginsForDB[index].ID = logins[index].ID
		loginsForDB[index].Name = logins[index].Name
		loginsForDB[index].URI = logins[index].URI
		loginsForDB[index].Login = logins[index].Login
		loginsForDB[index].Password = logins[index].Password
		loginsForDB[index].UserID = userID
	}

	return r.db.Save(loginsForDB).Error
}

func (r *GophKeeperRepo) LoadLogins() []viewsets.LoginForList {
	userID := r.getUserID()
	var logins []models.Login
	r.db.Where("user_id", userID).Find(&logins)
	if len(logins) == 0 {
		return nil
	}

	loginsViewSet := make([]viewsets.LoginForList, len(logins))

	for index := range logins {
		loginsViewSet[index].ID = logins[index].ID
		loginsViewSet[index].Name = logins[index].Name
		loginsViewSet[index].URI = logins[index].URI
	}

	return loginsViewSet
}

func (r *GophKeeperRepo) GetLoginByID(loginID uuid.UUID) (login entity.Login, err error) {
	var loginFromDB models.Login
	if err = r.db.Find(&loginFromDB, loginID).Error; loginFromDB.ID == uuid.Nil || err != nil {
		return login, errLoginNotFound
	}

	login.ID = loginFromDB.ID
	login.Login = loginFromDB.Login
	login.Name = loginFromDB.Name
	login.Password = loginFromDB.Password
	login.URI = loginFromDB.URI

	return
}

func (r *GophKeeperRepo) DelLogin(loginID uuid.UUID) error {
	return r.db.Unscoped().Delete(&models.Login{}, loginID).Error
}
