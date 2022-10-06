package integration_test

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
	"github.com/dimk00z/GophKeeper/internal/entity"
)

const (
	// Attempts connection
	// host       = "app:8080".
	host       = "localhost:8080"
	healthPath = "http://" + host + "/api/v1/health"
	attempts   = 20

	// HTTP REST.
	basePath = "http://" + host + "/api/v1"
)

var testUser = entity.User{ //nolint:gochecknoglobals // test entity
	Email:    "test@test.tt",
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

// HTTP GET: /auth/logout.
func TestHTTPDoUserLogout(t *testing.T) {
	Test(t,
		Description("UserLogin Success"),
		Get(basePath+"/auth/logout"),
		Send().Headers("Content-Type").Add("application/json"),
		Expect().Status().Equal(http.StatusOK),
	)
}

// // RabbitMQ RPC Client: getHistory.
// func TestRMQClientRPC(t *testing.T) {
// 	rmqClient, err := client.New(rmqURL, rpcServerExchange, rpcClientExchange)
// 	if err != nil {
// 		t.Fatal("RabbitMQ RPC Client - init error - client.New")
// 	}

// 	defer func() {
// 		err = rmqClient.Shutdown()
// 		if err != nil {
// 			t.Fatal("RabbitMQ RPC Client - shutdown error - rmqClient.RemoteCall", err)
// 		}
// 	}()

// 	type GophKeeper struct {
// 		Source      string `json:"source"`
// 		Destination string `json:"destination"`
// 		Original    string `json:"original"`
// 		GophKeeper  string `json:"GophKeeper"`
// 	}

// 	type historyResponse struct {
// 		History []GophKeeper `json:"history"`
// 	}

// 	for i := 0; i < requests; i++ {
// 		var history historyResponse

// 		err = rmqClient.RemoteCall("getHistory", nil, &history)
// 		if err != nil {
// 			t.Fatal("RabbitMQ RPC Client - remote call error - rmqClient.RemoteCall", err)
// 		}

// 		if history.History[0].Original != "текст для перевода" {
// 			t.Fatal("Original != текст для перевода")
// 		}
// 	}
// }
