package service

import (
	"Faceit/src/internal/model"
	"Faceit/src/mocks"
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type mocksUserService struct {
	nonRelationalUserDBRepository *mocks.MockNonRelationalUserDBRepository
}

func TestCreateUser(t *testing.T) {
	// · Mocks · //
	userMock := model.User{
		FirstName: "Niels",
		LastName:  "Sanchez",
		Nickname:  "Raws",
		Password:  "Niels1",
		Email:     "niels@niels.com",
		Country:   "SP",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// · Tests · //
	type want struct {
		result *model.User
		err    error
	}

	tests := []struct {
		name   string
		user   model.User
		want   want
		result string
		mocks  func(m mocksUserService)
	}{
		{
			name: "Should create user succesfully",
			user: userMock,
			want: want{
				result: &userMock,
				err:    nil,
			},
			mocks: func(m mocksUserService) {
				m.nonRelationalUserDBRepository.EXPECT().CreateUser(context.Background(), &userMock).Return(&userMock, nil)
			},
		},
	}

	// · Runner · //
	for _, tt := range tests {
		tt := tt

		// Prepare
		m := mocksUserService{
			nonRelationalUserDBRepository: mocks.NewMockNonRelationalUserDBRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		userService := NewUserService(m.nonRelationalUserDBRepository)

		// Execute
		result, err := userService.CreateUser(context.Background(), tt.user)

		// Verify
		if tt.want.err != nil && err != nil {
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}
		assert.Equal(t, &tt.want.result, &result)
	}
}
