package clientapi

import (
	"github.com/dimk00z/GophKeeper/internal/entity"
)

const loginsEndpoint = "api/v1/user/logins"

func (api *GophKeeperClientAPI) GetLogins(accessToken string) (logins []entity.Login, err error) {
	if err := api.getEntities(&logins, accessToken, loginsEndpoint); err != nil {
		return nil, err
	}

	return logins, nil
}

func (api *GophKeeperClientAPI) AddLogin(accessToken string, login *entity.Login) error {
	return api.addEntity(login, accessToken, loginsEndpoint)
}
