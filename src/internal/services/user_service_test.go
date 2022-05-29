package service

import (
	"Faceit/src/internal/model"
	"Faceit/src/mocks"
	"context"
	"errors"
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
		{
			name: "Should return error - Failed to query",
			user: userMock,
			want: want{
				result: nil,
				err:    errors.New("Failed to query"),
			},
			mocks: func(m mocksUserService) {
				m.nonRelationalUserDBRepository.EXPECT().CreateUser(context.Background(), &userMock).Return(nil, errors.New("Failed to query"))
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

func TestUpdateUser(t *testing.T) {
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
	userId := ""

	// · Tests · //
	type want struct {
		result *model.User
		err    error
	}

	tests := []struct {
		name   string
		user   model.User
		userId string
		want   want
		result string
		mocks  func(m mocksUserService)
	}{
		{
			name:   "Should update user succesfully",
			user:   userMock,
			userId: userId,
			want: want{
				result: &userMock,
				err:    nil,
			},
			mocks: func(m mocksUserService) {
				m.nonRelationalUserDBRepository.EXPECT().UpdateUser(context.Background(), userId, &userMock).Return(&userMock, nil)
			},
		},
		{
			name: "Should return error - userId not found",
			user: userMock,
			want: want{
				result: nil,
				err:    errors.New("userId not found"),
			},
			mocks: func(m mocksUserService) {
				m.nonRelationalUserDBRepository.EXPECT().UpdateUser(context.Background(), userId, &userMock).Return(nil, errors.New("userId not found"))
			},
		},
		{
			name: "Should return error - Failed to query",
			user: userMock,
			want: want{
				result: nil,
				err:    errors.New("Failed to query"),
			},
			mocks: func(m mocksUserService) {
				m.nonRelationalUserDBRepository.EXPECT().UpdateUser(context.Background(), userId, &userMock).Return(nil, errors.New("Failed to query"))
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
		result, err := userService.UpdateUser(context.Background(), tt.userId, tt.user)

		// Verify
		if tt.want.err != nil && err != nil {
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}
		assert.Equal(t, &tt.want.result, &result)
	}
}

func TestDeleteUser(t *testing.T) {
	// · Mocks · //
	userId := ""

	// · Tests · //
	type want struct {
		result *model.User
		err    error
	}

	tests := []struct {
		name   string
		user   model.User
		userId string
		want   want
		result string
		mocks  func(m mocksUserService)
	}{
		{
			name:   "Should update user succesfully",
			userId: userId,
			want: want{
				err: nil,
			},
			mocks: func(m mocksUserService) {
				m.nonRelationalUserDBRepository.EXPECT().DeleteUser(context.Background(), userId).Return(nil)
			},
		},
		{
			name: "Should return error - userId not found",
			want: want{
				result: nil,
				err:    errors.New("userId not found"),
			},
			mocks: func(m mocksUserService) {
				m.nonRelationalUserDBRepository.EXPECT().DeleteUser(context.Background(), userId).Return(errors.New("userId not found"))
			},
		},
		{
			name: "Should return error - Failed to query",
			want: want{
				result: nil,
				err:    errors.New("Failed to query"),
			},
			mocks: func(m mocksUserService) {
				m.nonRelationalUserDBRepository.EXPECT().DeleteUser(context.Background(), userId).Return(errors.New("Failed to query"))
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
		err := userService.DeleteUser(context.Background(), tt.userId)

		// Verify
		if tt.want.err != nil && err != nil {
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}
	}
}

// TODO
func TestGetUsers(t *testing.T) {
	// · Mocks · //
	userId := ""

	// · Tests · //
	type want struct {
		result *model.User
		err    error
	}

	tests := []struct {
		name   string
		user   model.User
		userId string
		want   want
		result string
		mocks  func(m mocksUserService)
	}{
		{
			name:   "Should update user succesfully",
			userId: userId,
			want: want{
				err: nil,
			},
			mocks: func(m mocksUserService) {
				m.nonRelationalUserDBRepository.EXPECT().DeleteUser(context.Background(), userId).Return(nil)
			},
		},
		{
			name: "Should return error - userId not found",
			want: want{
				result: nil,
				err:    errors.New("userId not found"),
			},
			mocks: func(m mocksUserService) {
				m.nonRelationalUserDBRepository.EXPECT().DeleteUser(context.Background(), userId).Return(errors.New("userId not found"))
			},
		},
		{
			name: "Should return error - Failed to query",
			want: want{
				result: nil,
				err:    errors.New("Failed to query"),
			},
			mocks: func(m mocksUserService) {
				m.nonRelationalUserDBRepository.EXPECT().DeleteUser(context.Background(), userId).Return(errors.New("Failed to query"))
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
		err := userService.DeleteUser(context.Background(), tt.userId)

		// Verify
		if tt.want.err != nil && err != nil {
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}
	}
}
