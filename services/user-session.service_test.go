package services_test

import (
	"go-demo/model"
	"go-demo/repositories"
	"go-demo/services"
	"testing"

	"github.com/alicebob/miniredis/v2"
	redis "github.com/go-redis/redis/v8"
)

func BeforeTest(t *testing.T, md bool) (services.UserService, *miniredis.Miniredis, model.SessionData) {

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	userSsRepo := repositories.NewUsersRepository(client)

	mockSession := model.SessionData{
		Token:        "token_110010101101",
		RefreshToken: "refreshtoken_1111000111001",
		Fingerprint:  "finger123",
		CoreId:       "1212",
		FirstName:    "felipe",
		LastName:     "elgueta",
		Country:      "colbun",
		Client:       "client123",
		Ttl:          "ttl123",
	}

	return services.NewUserService(userSsRepo, md), s, mockSession
}

func TestCreateUserSession(t *testing.T) {

	userSvc, mr, mockSession := BeforeTest(t, true)
	userID := "pepito123"

	err := userSvc.CreateUserSession(userID, "client", mockSession, "lala")

	if err != nil {
		t.Errorf("FAIL err: expected nil but got %s ", err)
	}

	t.Run("err = nil", func(t *testing.T) {

		mr.Set(userID, "anything")
		err := userSvc.CreateUserSession(userID, "client", mockSession, "lala")

		if err != nil {
			t.Errorf("FAIL err: expected nil but got %s ", err)
		}

	})

	t.Run("multiDevice  false", func(t *testing.T) {

		mr := miniredis.RunT(t)
		client := redis.NewClient(&redis.Options{
			Addr: mr.Addr(),
		})

		userSsRepo := repositories.NewUsersRepository(client)
		svc := services.NewUserService(userSsRepo, false)

		mr.Set(userID, "anything")
		err := svc.CreateUserSession(userID, "client", mockSession, "lala")

		if err != nil {
			t.Errorf("FAIL err: expected nil but got %s ", err)
		}

	})

}

func TestGetUserSessions(t *testing.T) {

	userSvc, _, mockSession := BeforeTest(t, true)

	keyID := "pepe123"
	userSvc.CreateUserSession(keyID, "client", mockSession, "lala")
	_, err := userSvc.GetUserSessions(keyID)

	if err != nil {
		t.Errorf("FAIL err: expected nil but got %s ", err)
	}
}
