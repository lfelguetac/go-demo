package services

import (
	"go-demo/model"
	"go-demo/repositories"
	"go-demo/utils"
)

type UserService interface {
	CreateUserSession(userId, client string, session model.SessionData, ttl string) error
	GetUserSessions(userId string) (*[]model.SessionData, error)
}

type service struct {
	repo repositories.UserSessionRepository
}

var multiDevice bool

func NewUserService(repository repositories.UserSessionRepository, md bool) UserService {
	multiDevice = md
	return &service{repo: repository}
}

func (svc *service) CreateUserSession(userId string, client string, session model.SessionData, ttl string) error {
	userSession, _err := svc.repo.GetUserSessions(userId)

	// TODO: check only KEY not found error
	if _err != nil {
		userSession := model.UserSession{
			Sessions: []model.SessionData{session},
		}
		_err = svc.repo.SetUserSession(userId, userSession, ttl)
	} else {
		if multiDevice {

			sessions := utils.DeleteFirstClient(userSession.Sessions, client)
			userSession.Sessions = append([]model.SessionData{session}, sessions...)
			_err = svc.repo.SetUserSession(userId, *userSession, ttl)
		} else {
			userSession.Sessions = []model.SessionData{session}
			_err = svc.repo.SetUserSession(userId, *userSession, ttl)
		}
	}

	if _err != nil {
		return _err
	}
	return nil
}

func (svc *service) GetUserSessions(userId string) (*[]model.SessionData, error) {
	userSession, _err := svc.repo.GetUserSessions(userId)
	if _err != nil {

		return nil, _err
	}
	return &userSession.Sessions, nil
}
