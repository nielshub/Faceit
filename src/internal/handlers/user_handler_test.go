package handlers

import (
	"Faceit/src/internal/model"
	"Faceit/src/internal/ports"
	service "Faceit/src/internal/services"
	"Faceit/src/mocks"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockUserHandler struct {
	router           *gin.RouterGroup
	userService      ports.UserService
	publisherService *mocks.MockPublisherService
}

type mocksUserService struct {
	nonRelationalUserDBRepository *mocks.MockNonRelationalUserDBRepository
}

func TestUserDelete(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// · Mocks · //
	id := "userId"
	response := "User:" + id + " has been deleted properly."
	outMsg := model.Message{
		Queue:       "",
		ContentType: "text/plain",
		Data:        []byte(response),
	}
	// · Tests · //
	type want struct {
		code     int
		response string
		err      error
	}

	tests := []struct {
		name   string
		user   model.User
		url    string
		want   want
		result string
		mocks  func(mUS mocksUserService, mPS mockUserHandler)
	}{
		{
			name: "Should delete user succesfully",
			url:  "/user/delete/" + id,
			want: want{
				code:     http.StatusOK,
				response: "\"User:userId has been deleted properly.\"",
				err:      nil,
			},
			mocks: func(mUS mocksUserService, mPS mockUserHandler) {
				mUS.nonRelationalUserDBRepository.EXPECT().DeleteUser(context.Background(), id).Return(nil)
				mPS.publisherService.EXPECT().Publish(outMsg).Return(nil)
			},
		},
		{
			name: "Should return error - Failed to query DB",
			url:  "/user/delete/" + id,
			want: want{
				code: http.StatusInternalServerError,
				response: `{
					"message": "Error deleting user"
				}`,
				err: errors.New("Error deleting user"),
			},
			mocks: func(mUS mocksUserService, mPS mockUserHandler) {
				mUS.nonRelationalUserDBRepository.EXPECT().DeleteUser(context.Background(), id).Return(errors.New("Error deleting user"))
			},
		},
		{
			name: "Should return error - Failed to publish to rabbitMQ",
			url:  "/user/delete/" + id,
			want: want{
				code: http.StatusInternalServerError,
				response: `{
					"message": "User:userId has been deleted properly. Message has not been sent to rabbitMQ."
				}`,
				err: errors.New("Error deleting user"),
			},
			mocks: func(mUS mocksUserService, mPS mockUserHandler) {
				mUS.nonRelationalUserDBRepository.EXPECT().DeleteUser(context.Background(), id).Return(nil)
				mPS.publisherService.EXPECT().Publish(outMsg).Return(errors.New(""))
			},
		},
	}

	// · Runner · //
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// Prepare
			mUS := mocksUserService{
				nonRelationalUserDBRepository: mocks.NewMockNonRelationalUserDBRepository(gomock.NewController(t)),
			}
			w := httptest.NewRecorder()
			r := gin.Default()
			app := r.Group("/")

			mPS := mockUserHandler{
				router:           app,
				userService:      service.NewUserService(mUS.nonRelationalUserDBRepository),
				publisherService: mocks.NewMockPublisherService(gomock.NewController(t)),
			}

			tt.mocks(mUS, mPS)
			NewUserHandler(mPS.router, mPS.userService, mPS.publisherService)

			req, err := http.NewRequest("POST", tt.url, bytes.NewBufferString(""))
			require.NoError(t, err)
			r.ServeHTTP(w, req)
			assert.JSONEq(t, tt.want.response, w.Body.String())
			assert.Equal(t, tt.want.code, w.Code)
		})

	}

}

func TestGetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// · Mocks · //
	userMock := model.User{
		FirstName: "Niels",
		LastName:  "Sanchez",
		Nickname:  "Raws",
		Password:  "Niels1",
		Email:     "niels@niels.com",
		Country:   "SP",
		CreatedAt: time.Date(2022, time.May, 30, 15, 37, 41, 742045900, time.Local),
		UpdatedAt: time.Date(2022, time.May, 30, 15, 37, 41, 742045900, time.Local),
	}
	userMockArray := []model.User{
		userMock,
		userMock,
	}
	response, _ := json.Marshal(userMockArray)
	key := ""
	value := ""
	// · Tests · //
	type want struct {
		code     int
		response string
		err      error
	}

	tests := []struct {
		name   string
		url    string
		want   want
		result string
		mocks  func(mUS mocksUserService, mPS mockUserHandler)
	}{
		{
			name: "Should get users succesfully",
			url:  "/user/get?key=" + key + "&value=" + value,
			want: want{
				code:     http.StatusOK,
				response: string(response),
				err:      nil,
			},
			mocks: func(mUS mocksUserService, mPS mockUserHandler) {
				mUS.nonRelationalUserDBRepository.EXPECT().GetAllUsers(context.Background()).Return(userMockArray, nil)
			},
		},
		{
			name: "Should return error - Failed to query DB",
			url:  "/user/get?key=" + key + "&value=" + value,
			want: want{
				code: http.StatusInternalServerError,
				response: `{
					"message": "Error getting users"
				}`,
				err: errors.New("Error getting users"),
			},
			mocks: func(mUS mocksUserService, mPS mockUserHandler) {
				mUS.nonRelationalUserDBRepository.EXPECT().GetAllUsers(context.Background()).Return(nil, errors.New("DB server error"))
			},
		},
	}

	// · Runner · //
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// Prepare
			mUS := mocksUserService{
				nonRelationalUserDBRepository: mocks.NewMockNonRelationalUserDBRepository(gomock.NewController(t)),
			}
			w := httptest.NewRecorder()
			r := gin.Default()
			app := r.Group("/")

			mPS := mockUserHandler{
				router:           app,
				userService:      service.NewUserService(mUS.nonRelationalUserDBRepository),
				publisherService: mocks.NewMockPublisherService(gomock.NewController(t)),
			}

			tt.mocks(mUS, mPS)
			NewUserHandler(mPS.router, mPS.userService, mPS.publisherService)

			req, err := http.NewRequest("GET", tt.url, bytes.NewBufferString(""))
			require.NoError(t, err)
			r.ServeHTTP(w, req)
			assert.JSONEq(t, tt.want.response, w.Body.String())
			assert.Equal(t, tt.want.code, w.Code)
		})

	}

}
