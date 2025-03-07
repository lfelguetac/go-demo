package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"go-demo/logger"
	"go-demo/model"
	"time"

	// "github.com/go-redis/redis"
	redis "github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var log *logger.FpayLogger = logger.GetLogger()

type UserSessionRepository interface {
	SetUserSession(userId string, userSession model.UserSession, ttl string) error
	GetUserSessions(userId string) (*model.UserSession, error)
}

type db struct {
	dbConnection *redis.Client
}

func NewUsersRepository(redisC *redis.Client) UserSessionRepository {
	return &db{
		dbConnection: redisC,
	}
}

func (r *db) SetUserSession(userId string, userSession model.UserSession, ttl string) error {
	us, _ := json.Marshal(userSession)
	_err := r.dbConnection.Set(ctx, userId, us, GetTtlTime(ttl)).Err()

	if _err != nil {
		log.Error(_err.Error())
		return errors.New(_err.Error())
	}
	return nil
}

func (r *db) GetUserSessions(userId string) (*model.UserSession, error) {
	result, _err := r.dbConnection.Get(ctx, userId).Bytes()

	var userSession model.UserSession
	json.Unmarshal(result, &userSession)
	if _err != nil {
		log.Error(_err.Error())
		return nil, errors.New("id not found")
	}
	return &userSession, nil
}

func GetTtlTime(ttl string) time.Duration {
	ttlHour, _err := time.ParseDuration(ttl)
	if _err != nil {
		return time.Hour
	}
	return ttlHour
}
