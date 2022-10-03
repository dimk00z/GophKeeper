package usecase

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/utils/errs"
	"github.com/go-resty/resty/v2"
)

func (uc *GophKeeperClientUseCase) Login(user *entity.User) {
	client := resty.New()
	var token entity.JWT
	body := fmt.Sprintf(`{"email":%q, "password":%q}`, user.Email, user.Password)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&token).
		Post(fmt.Sprintf("%s/api/v1/auth/login", uc.cfg.Server.URL))
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode() == http.StatusBadRequest || resp.StatusCode() == http.StatusInternalServerError {
		color.Red("Server error: %s", errs.ParseServerError(resp.Body()))

		return
	}

	if !uc.repo.UserExistsByEmail(user.Email) {
		err = uc.repo.AddUser(user)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err = uc.repo.UpdateUserToken(user, &token); err != nil {
		log.Fatal(err)
	}
	color.Green("Got authorisation token for %q", user.Email)
}

func (uc *GophKeeperClientUseCase) Register(user *entity.User) {
	client := resty.New()
	body := fmt.Sprintf(`{"email":%q, "password":%q}`, user.Email, user.Password)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(user).
		Post(fmt.Sprintf("%s/api/v1/auth/register", uc.cfg.Server.URL))
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode() == http.StatusBadRequest || resp.StatusCode() == http.StatusInternalServerError {
		color.Red("Server error: %s", errs.ParseServerError(resp.Body()))

		return
	}
	uc.repo.AddUser(user)
	color.Green("User registered")
	color.Green("ID: %v", user.ID)
	color.Green("ID: %s", user.Email)
}
