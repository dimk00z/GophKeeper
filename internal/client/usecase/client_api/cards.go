package clientapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/utils/errs"
	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
	"golang.org/x/exp/slices"
)

func (api *GophKeeperClientAPI) GetCards(accessToken string) (cards []entity.Card, err error) {
	client := resty.New()
	client.SetAuthToken(accessToken)
	resp, err := client.R().
		SetResult(&cards).
		Get(fmt.Sprintf("%s/api/v1/user/cards", api.serverURL))
	if err != nil {
		log.Println(err)

		return
	}

	badCodes := []int{http.StatusBadRequest, http.StatusInternalServerError, http.StatusUnauthorized}
	if slices.Contains(badCodes, resp.StatusCode()) {
		errMessage := errs.ParseServerError(resp.Body())
		color.Red("Server error: %s", errMessage)
		err = errServer

		return
	}

	return
}

func (api *GophKeeperClientAPI) AddCard(accessToken string, card *entity.Card) error {
	client := resty.New()
	client.SetAuthToken(accessToken)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(card).
		SetResult(card).
		Post(fmt.Sprintf("%s/api/v1/user/cards", api.serverURL))
	if err != nil {
		log.Fatal(err)
	}

	badCodes := []int{http.StatusBadRequest, http.StatusInternalServerError, http.StatusUnauthorized}
	if slices.Contains(badCodes, resp.StatusCode()) {
		errMessage := errs.ParseServerError(resp.Body())
		color.Red("Server error: %s", errMessage)

		return errServer
	}

	return nil
}
