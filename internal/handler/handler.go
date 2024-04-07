package handler

import (
	"bannerService/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Route(e *echo.Echo) {
	//userBanner := e.Group("/user_banner")
	//userBanner.GET("/", h.getBanner)

	banner := e.Group("/banner")
	//banner.GET("/", h.filterBanners)
	banner.POST("/", h.createBanner)
	//banner.PATCH("/:id", h.updateBanner)
	//banner.DELETE("/:id", h.deleteBanner)

	//e.GET("/swagger/*", echoSwagger.WrapHandler)
}
