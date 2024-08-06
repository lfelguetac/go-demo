package controller_test

import (
	"bytes"
	"encoding/json"
	"go-demo/delivery/controller"
	"go-demo/model"
	"go-demo/repositories"
	"go-demo/services"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"

	redis "github.com/go-redis/redis/v8"
)

func BeforeTest(t *testing.T) (controller.UserSessionController, *miniredis.Miniredis, model.SessionRequest, string) {

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	keyID := "pepe123"

	reqBody := model.SessionRequest{
		UserId: keyID,
		Client: "client123",
		Ttl:    "ttl123",
		Data: model.SessionData{
			Token:        "token_1110110010111",
			RefreshToken: "rtokeen_1100011001",
			Fingerprint:  "finger123",
			CoreId:       "coreId123",
			FirstName:    "felipe",
			LastName:     "elgueta",
			Country:      "colbun",
			Client:       "client123",
			Ttl:          "ttl123",
		},
	}

	userRepository := repositories.NewUsersRepository(client)
	userService := services.NewUserService(userRepository, true)
	userController := controller.NewUserSessionController(userService)
	return userController, s, reqBody, keyID
}

func TestCreateUserSession(t *testing.T) {

	ctrl, _, reqBody, _ := BeforeTest(t)
	w := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(w)

	t.Run("ShouldBindJSON error", func(t *testing.T) {

		payload, _ := json.Marshal(reqBody)
		testContext.Request, _ = http.NewRequest("POST", "/user", bytes.NewBuffer(payload))
		testContext.Request.Header.Set("Content-Type", "application/json")

		ctrl.CreateUserSession(testContext)

		expected := 201
		got := testContext.Writer.Status()

		if expected != got {
			t.Errorf("expected %d but got %d", expected, got)
		}

	})

}

func TestGetUserSessions(t *testing.T) {

	ctrl, mr, _, keyID := BeforeTest(t)

	mr.Set(keyID, "anything")
	str, _ := mr.Get(keyID)
	t.Logf("keyID: %s", str)

	testContextGet, _ := gin.CreateTestContext(httptest.NewRecorder())
	testContextGet.Request, _ = http.NewRequest("GET", "/user/pepe123", nil)
	testContextGet.Params = []gin.Param{
		{
			Key:   "userId",
			Value: keyID,
		},
	}
	ctrl.GetUserSessions(testContextGet)

	expected := 200
	got := testContextGet.Writer.Status()

	if expected != got {
		t.Errorf("expected 200 but got %d", got)
	}

}
