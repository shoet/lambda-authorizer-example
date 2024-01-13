package service

import (
	"context"
	"errors"

	"github.com/shoet/lambda-authorizer-example/internal/infrastracture"
)

type KeyValueStore interface {
	Get(key string) (interface{}, error)
}

type AuthService struct {
	KVS KeyValueStore
}

func NewAuthService(kvs KeyValueStore) *AuthService {
	return &AuthService{
		KVS: kvs,
	}
}

func (a *AuthService) Login(ctx context.Context, email, password string) (bool, error) {
	value, err := a.KVS.Get(email)
	if err != nil {
		if errors.Is(err, infrastracture.ErrEntityNotFound) {
			return false, nil
		}
		return false, err
	}
	if value != password {
		return false, nil
	}
	return true, nil
}

func (a *AuthService) GenerateToken(ctx context.Context) (string, error) {
	return "XXXXXDUMMYTOKENXXXXX", nil
}
