package handlers

import (
	"Faceit/src/internal/model"
	"Faceit/src/internal/ports"
	service "Faceit/src/internal/services"
	"Faceit/src/mocks"
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

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
	jsonResponse := "\"User:userId has been deleted properly.\""
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
				code:     200,
				response: jsonResponse,
				err:      nil,
			},
			mocks: func(mUS mocksUserService, mPS mockUserHandler) {
				mUS.nonRelationalUserDBRepository.EXPECT().DeleteUser(context.Background(), id).Return(nil)
				mPS.publisherService.EXPECT().Publish(outMsg).Return(nil)
			},
		},
		// {
		// 	name: "Should return error - Failed to query DB",
		// 	url:  "/user/delete/" + id,
		// 	want: want{
		// 		code:     500,
		// 		response: "Internal server error",
		// 		err:      errors.New("Failed to query"),
		// 	},
		// 	mocks: func(mUS mocksUserService, mPS mockUserHandler) {
		// 		mUS.nonRelationalUserDBRepository.EXPECT().DeleteUser(context.Background(), id).Return(nil)
		// 		//mPS.publisherService.EXPECT().Publish(outMsg).Return(nil)
		// 	},
		// },
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
