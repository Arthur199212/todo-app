package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"
	"todo-app/models"
	"todo-app/pkg/service"
	mock_service "todo-app/pkg/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockAuthorization, user models.User)

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"email":"test@test.com","password":"!Welcome01"}`,
			inputUser: models.User{
				Email:    "test@test.com",
				Password: "!Welcome01",
			},
			mockBehaviour: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Not valid email",
			inputBody:            `{"email":"test#test.com","password":"!Welcome01"}`,
			mockBehaviour:        func(s *mock_service.MockAuthorization, user models.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"email: must be a valid email address."}`,
		},
		{
			name:      "Service failure",
			inputBody: `{"email":"test@test.com","password":"!Welcome01"}`,
			inputUser: models.User{
				Email:    "test@test.com",
				Password: "!Welcome01",
			},
			mockBehaviour: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().CreateUser(user).Return(0, errors.New("some service failure"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"some service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// init deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehaviour(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			// test
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			// test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			// perform request
			r.ServeHTTP(w, req)

			// assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
