package handler

import (
	"bannerService/internal/entity"
	"bannerService/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (h *Handler) getBannerHistory(ctx echo.Context) error {
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

	history, err := h.services.GetBannerHistory(ctx.Request().Context(), bannerId.BannerId)
	if err != nil {
		utils.Logger.Error("get banner history error", zap.String("error", err.Error()))
		return responseErr(err)
	}

	return responseOk(ctx, history)
}
