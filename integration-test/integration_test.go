package integration_test

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dimk00z/GophKeeper/internal/entity"
)

const (
	host = "app:8080"
	// host       = "localhost:8080" // for local testing
	healthPath = "http://" + host + "/api/v1/health"
	attempts   = 20

	// HTTP REST.
	basePath      = "http://" + host + "/api/v1"
	numberOfTests = 3
)

var testUser = entity.User{ //nolint:gochecknoglobals // test entity
	Email:    gofakeit.Email(),
	Password: "password",
}
var testUserToken = entity.JWT{} //nolint:gochecknoglobals // test entity

func TestMain(m *testing.M) {
	err := healthCheck(attempts)
	if err != nil {
		log.Fatalf("Integration tests: host %s is not available: %s", host, err)
	}

	log.Printf("Integration tests: host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

func healthCheck(attempts int) error {
	var err error

	for attempts > 0 {
		err = Do(Get(healthPath), Expect().Status().Equal(http.StatusOK))
		if err == nil {
			return nil
		}

		log.Printf("Integration tests: url %s is not available, attempts left: %d", healthPath, attempts)

		time.Sleep(time.Second)

		attempts--
	}

	return err
}

// HTTP POST: /auth/register.
func TestHTTPDoUserRegister(t *testing.T) {
	body := fmt.Sprintf(`{
		"email": %q,
		"password": %q
	}`, testUser.Email, testUser.Password)
	Test(t,
		Description("UserRegister Success"),
		Post(basePath+"/auth/register"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusCreated),
		Expect().Body().JSON().Contains("email", "uuid"),
	)

	Test(t,
		Description("UserRegister secod try with the same user data"),
		Post(basePath+"/auth/register"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().Contains("error"),
	)

	body = `{
		"email": "wrong_email",
		"password": 1223
	}`
	Test(t,
		Description("UserRegister Wrong users data"),
		Post(basePath+"/auth/register"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().Contains("error"),
	)
}

// HTTP POST: /auth/login.
func TestHTTPDoUserLogin(t *testing.T) {
	body := fmt.Sprintf(`{
		"email": %q,
		"password": %q
	}`, testUser.Email, testUser.Password)
	Test(t,
		Description("UserLogin Success"),
		Post(basePath+"/auth/login"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Contains("access_token", "refresh_token"),
		Store().Response().Body().JSON().JQ(".access_token").In(&testUserToken.AccessToken),
		Store().Response().Body().JSON().JQ(".refresh_token").In(&testUserToken.RefreshToken),
	)
	body = `{
		"email": "wrong@email.com",
		"password": "wrong_pass"
	}`
	Test(t,
		Description("UserRegister Wrong users data"),
		Post(basePath+"/auth/login"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().Contains("error"),
	)
}

func getTestCard() entity.Card {
	return entity.Card{
		Name:            gofakeit.NounAbstract(),
		CardHolderName:  gofakeit.LastName() + " " + gofakeit.Name(),
		Number:          gofakeit.CreditCardNumber(&gofakeit.CreditCardOptions{Gaps: true}),
		ExpirationMonth: "02",
		ExpirationYear:  "2022",
		Brand:           gofakeit.CreditCardType(),
		SecurityCode:    gofakeit.CreditCardCvv(),
	}
}

var testCards []entity.Card //nolint:gochecknoglobals // use for tests

// HTTP Post: /users/cards.
func TestHTTPAddUserCard(t *testing.T) {
	testCards = make([]entity.Card, numberOfTests)
	for i := 0; i < numberOfTests; i++ {
		testCards[i] = getTestCard()
	}
	Test(t,
		Description("UserLogin Add card without token"),
		Post(basePath+"/user/cards"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().JSON(&testCards[0]),
		Expect().Status().Equal(http.StatusUnauthorized),
		Expect().Body().JSON().Contains("error"),
	)
	for i := 0; i < numberOfTests; i++ {
		Test(t,
			Description("UserLogin Add card with token"),
			Post(basePath+"/user/cards"),
			Send().Headers("Content-Type").Add("application/json"),
			Send().Headers("Authorization").Add("Bearer "+testUserToken.AccessToken),
			Send().Body().JSON(&testCards[i]),
			Expect().Status().Equal(http.StatusAccepted),
			Store().Response().Body().JSON().JQ(".uuid").In(&testCards[i].ID),
		)
	}
	Test(t,
		Description("UserLogin Add card - empty body"),
		Post(basePath+"/user/cards"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Headers("Authorization").Add("Bearer "+testUserToken.AccessToken),
		Expect().Status().Equal(http.StatusBadRequest),
	)
}

// HTTP get: /users/cards.
func TestHTTPGetUserCard(t *testing.T) {
	Test(t,
		Description("UserLogin Get card without token"),
		Get(basePath+"/user/cards"),
		Send().Headers("Content-Type").Add("application/json"),
		Expect().Status().Equal(http.StatusUnauthorized),
		Expect().Body().JSON().Contains("error"),
	)

	var cards []entity.Card

	Test(t,
		Description("UserLogin Get card with token"),
		Get(basePath+"/user/cards"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Headers("Authorization").Add("Bearer "+testUserToken.AccessToken),
		Expect().Status().Equal(http.StatusOK),
		Store().Response().Body().JSON().In(&cards),
	)
	if len(cards) != numberOfTests {
		t.Errorf("Expected %v, got %v", numberOfTests, len(cards))
	}
}

// HTTP delete: /users/cards/:id.
func TestHTTPDelUserCard(t *testing.T) {
	Test(t,
		Description("UserLogin Del card"),
		Delete(basePath+"/user/cards/"+testCards[0].ID.String()),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Headers("Authorization").Add("Bearer "+testUserToken.AccessToken),
		Expect().Status().Equal(http.StatusAccepted),
	)

	var testCards []entity.Card

	Test(t,
		Description("UserLogin Get card after delete"),
		Get(basePath+"/user/cards"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Headers("Authorization").Add("Bearer "+testUserToken.AccessToken),
		Expect().Status().Equal(http.StatusOK),
		Store().Response().Body().JSON().In(&testCards),
	)
	if len(testCards) != numberOfTests-1 {
		t.Errorf("Expected %v, got %v", numberOfTests-1, len(testCards))
	}
}

// HTTP GET: /auth/logout.
func TestHTTPDoUserLogout(t *testing.T) {
	Test(t,
		Description("UserLogin Success"),
		Get(basePath+"/auth/logout"),
		Send().Headers("Content-Type").Add("application/json"),
		Expect().Status().Equal(http.StatusOK),
	)
}
