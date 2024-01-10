package service_test

import (
	"context"
	"testing"

	"github.com/shoet/lambda-authorizer-example/internal/infrastracture"
	"github.com/shoet/lambda-authorizer-example/internal/service"
	"github.com/stretchr/testify/mock"
)

type MockKeyValueStore struct {
	mock.Mock
}

func (k *MockKeyValueStore) Get(key string) (interface{}, error) {
	args := k.Called(key)
	return args.Get(0), args.Error(1)
}

func Test_AuthService_Login(t *testing.T) {
	type args struct {
		username string
		password string
		mockKvs  func() *MockKeyValueStore
	}
	type wants struct {
		ok bool
	}
	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name: "success",
			args: args{
				username: "username",
				password: "password",
				mockKvs: func() *MockKeyValueStore {
					mockKvs := new(MockKeyValueStore)
					mockKvs.On("Get", "username").Return("password", nil)
					return mockKvs
				},
			},
			wants: wants{
				ok: true,
			},
		},
		{
			name: "failed unmatched password",
			args: args{
				username: "username",
				password: "password",
				mockKvs: func() *MockKeyValueStore {
					mockKvs := new(MockKeyValueStore)
					mockKvs.On("Get", "username").Return("pass", nil)
					return mockKvs
				},
			},
			wants: wants{
				ok: false,
			},
		},
		{
			name: "failed user not found",
			args: args{
				username: "username",
				password: "password",
				mockKvs: func() *MockKeyValueStore {
					mockKvs := new(MockKeyValueStore)
					mockKvs.On("Get", "username").Return(nil, infrastracture.ErrEntityNotFound)
					return mockKvs
				},
			},
			wants: wants{
				ok: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockKvs := tt.args.mockKvs()

			authServce := service.NewAuthService(mockKvs)

			ok, err := authServce.Login(context.Background(), tt.args.username, tt.args.password)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.wants.ok != ok {
				t.Errorf("want %v, but %v", tt.wants.ok, ok)
			}
		})
	}
}
