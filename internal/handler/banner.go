package handler

import (
	"bannerService/internal/entity"
	"bannerService/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (h *Handler) createBanner(ctx echo.Context) error {
	err := h.checkAdminAuthorization(ctx.Request().Header.Get("Authorization"))
	if err != nil {
		utils.Logger.Error("incorrect auth data", zap.String("error", err.Error()))
		return err
	}

	var banner entity.Banner

	if err := ctx.Bind(&banner); err != nil {
		utils.Logger.Error("incorrect banner data", zap.String("error", err.Error()))
		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect banner data"))
	}

	utils.Logger.Debug("got banner to create", zap.Any("banner", banner))

	id, err := h.services.CreateBanner(ctx.Request().Context(), banner)
	if err != nil {
		utils.Logger.Error("banner creation error", zap.String("error", err.Error()))
		return responseErr(err)
	}

	utils.Logger.Info("banner created")
	utils.Logger.Debug("banner created with id", zap.Int("id", id))

	return responseCreated(ctx, ResponseId{BannerId: id})
}

func (h *Handler) deleteBanner(ctx echo.Context) error {
	err := h.checkAdminAuthorization(ctx.Request().Header.Get("Authorization"))
	if err != nil {
		utils.Logger.Error("incorrect auth data", zap.String("error", err.Error()))
		return err
	}

	var bannerId entity.BannerId

	if err := ctx.Bind(&bannerId); err != nil {
		utils.Logger.Error("incorrect data", zap.String("error", err.Error()))
		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect data"))
	}

	err = h.services.DeleteBanner(ctx.Request().Context(), bannerId.Id)
	if err != nil {
		utils.Logger.Error("banner deletion error", zap.String("error", err.Error()))
		return responseErr(err)
	}

	return responseDeleted(ctx, "banner was deleted")
}
