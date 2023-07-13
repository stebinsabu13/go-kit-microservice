package server

import (
	"context"
	"errors"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type Service interface {
	CreateUser(context.Context, string, string) error
	GetUser(context.Context, string) (User, error)
}
type service struct {
	logger log.Logger
}

func NewService(log log.Logger) Service {
	return &service{
		logger: log,
	}
}
func (s service) CreateUser(ctx context.Context, email string, password string) error {
	logger := log.With(s.logger, "method", "CreateUser")
	if email == "" || password == "" {
		level.Error(logger).Log("err", "invalid credentials")
		return errors.New("invalid credentials")
	}
	user[email] = User{
		Email:    email,
		Password: password,
	}
	logger.Log("create user", email)
	return nil
}

func (s service) GetUser(ctx context.Context, email string) (User, error) {
	logger := log.With(s.logger, "method", "GetUser")
	result := user[email]
	logger.Log("Get user", result)
	return result, nil
}
