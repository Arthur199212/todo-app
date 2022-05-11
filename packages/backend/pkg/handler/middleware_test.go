package handler

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"
	"todo-app/pkg/service"
	mock_service "todo-app/pkg/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_authRequired(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehaviour: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseUserIdFromToken(token).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "1",
		},
		{
			name:                 "No header",
			headerName:           "",
			mockBehaviour:        func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"not authorized"}`,
		},
		{
			name:                 "No token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			mockBehaviour:        func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name:        "Service failure",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehaviour: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseUserIdFromToken(token).Return(0, errors.New("failed to parse token"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid token"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// init deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehaviour(auth, testCase.token)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			// test server
			r := gin.New()
			r.POST("/protected", handler.authRequired, func(c *gin.Context) {
				id, _ := c.Get(userCtx)
				c.String(200, fmt.Sprintf("%d", id.(int)))
			})

			// test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			// make a request
			r.ServeHTTP(w, req)

			// assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
