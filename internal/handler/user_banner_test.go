package handler

import (
	"bannerService/internal/config"
	"bannerService/internal/entity"
	"bannerService/internal/service"
	mock_service "bannerService/internal/service/mocks"
	"bannerService/internal/utils"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_getBanner(t *testing.T) {
	type mockBehavior func(r *mock_service.MockUserBanner, ctx context.Context, banner entity.UserBanner)
	tests := []struct {
		name                 string
		token                string
		queryString          string
		inputBanner          entity.UserBanner
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "ok",
			token:       "user_test_token",
			queryString: "?tag_id=1&feature_id=1",
			inputBanner: entity.UserBanner{
				TagId:           1,
				FeatureId:       1,
				UseLastRevision: false,
			},
			mockBehavior: func(r *mock_service.MockUserBanner, ctx context.Context, banner entity.UserBanner) {
				r.EXPECT().GetBanner(ctx, banner).Return(map[string]interface{}{
					"text":  "some_text",
					"title": "some_title",
					"url":   "some_url",
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"text":"some_text","title":"some_title","url":"some_url"}` + "\n",
		},
		{
			name:        "missing params",
			token:       "user_test_token",
			queryString: "?tag_id=1",
			inputBanner: entity.UserBanner{
				TagId:           1,
				FeatureId:       1,
				UseLastRevision: false,
			},
			mockBehavior:         func(r *mock_service.MockUserBanner, ctx context.Context, banner entity.UserBanner) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"incorrect banner data, missing tag or feature: bad request"}` + "\n",
		},
		{
			name:        "no data",
			token:       "user_test_token",
			queryString: "?tag_id=1&feature_id=1",
			inputBanner: entity.UserBanner{
				TagId:           1,
				FeatureId:       1,
				UseLastRevision: false,
			},
			mockBehavior: func(r *mock_service.MockUserBanner, ctx context.Context, banner entity.UserBanner) {
				r.EXPECT().GetBanner(ctx, banner).Return(nil,
					errors.Wrap(utils.ErrNotFound,
						fmt.Sprintf("Banner with tag %v and feature %v not found", banner.TagId, banner.FeatureId)))
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"Banner with tag 1 and feature 1 not found: not found"}` + "\n",
		},
		{
			name:        "no token",
			token:       "",
			queryString: "?tag_id=1&feature_id=1",
			inputBanner: entity.UserBanner{
				TagId:           1,
				FeatureId:       1,
				UseLastRevision: false,
			},
			mockBehavior:         func(r *mock_service.MockUserBanner, ctx context.Context, banner entity.UserBanner) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"authorization required: no authorization"}` + "\n",
		},
		{
			name:        "wrong token",
			token:       "wrong_token",
			queryString: "?tag_id=1&feature_id=1",
			inputBanner: entity.UserBanner{
				TagId:           1,
				FeatureId:       1,
				UseLastRevision: false,
			},
			mockBehavior:         func(r *mock_service.MockUserBanner, ctx context.Context, banner entity.UserBanner) {},
			expectedStatusCode:   403,
			expectedResponseBody: `{"message":"access denied: access denied"}` + "\n",
		},
	}
	utils.CreateLogger()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockUserBanner(c)
			tt.mockBehavior(repo, context.Background(), tt.inputBanner)

			services := &service.Service{UserBanner: repo}
			handler := New(services, &config.Tokens{
				User:  "user_test_token",
				Admin: "admin_test_token",
			})

			r := echo.New()
			r.GET("/", handler.checkUserAuth(handler.getBanner))

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/"+tt.queryString, nil)
			req.Header.Set(echo.HeaderAuthorization, tt.token)

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())
		})
	}
}
