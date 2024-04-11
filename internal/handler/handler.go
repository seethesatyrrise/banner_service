package handler

import (
	"bannerService/internal/config"
	"bannerService/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services *service.Service
	tokens   *token
}

func New(services *service.Service, tokens *config.Tokens) *Handler {
	return &Handler{
		services: services,
		tokens: &token{
			userToken:  tokens.User,
			adminToken: tokens.Admin,
		},
	}
}

func (h *Handler) Route(e *echo.Echo) {
	userBanner := e.Group("/user_banner")
	userBanner.GET("/", h.getBanner)

	banner := e.Group("/banner")
	banner.GET("/", h.filterBanners)
	banner.POST("/", h.createBanner)
	banner.PATCH("/:banner_id", h.updateBanner)
	banner.DELETE("/:banner_id", h.deleteBanner)

	bannerHistory := e.Group("/banner_history")
	bannerHistory.GET("/:banner_id", h.getBannerHistory)
	//bannerHistory.POST("/:banner_id", h.selectBannerVersion)

	//e.GET("/swagger/*", echoSwagger.WrapHandler)
}
