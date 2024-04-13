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
	userBanner.GET("/", h.checkUserAuth(h.getBanner))

	banner := e.Group("/banner")
	banner.GET("/", h.checkAdminAuth(h.filterBanners))
	banner.POST("/", h.checkAdminAuth(h.createBanner))
	banner.PATCH("/:banner_id", h.checkAdminAuth(h.updateBanner))

	bannerHistory := e.Group("/banner_history")
	bannerHistory.GET("/:banner_id", h.checkAdminAuth(h.getBannerHistory))
	bannerHistory.POST("/:banner_id", h.checkAdminAuth(h.setBannerVersion))

	deletion := e.Group("/delete")
	deletion.DELETE("/id/:banner_id", h.checkAdminAuth(h.deleteId))
	deletion.DELETE("/feature/:feature_id", h.checkAdminAuth(h.deleteFeature))
	deletion.DELETE("/tag/:tag_id", h.checkAdminAuth(h.deleteTag))

	//e.GET("/swagger/*", echoSwagger.WrapHandler)
}
