package handler

import (
	"bannerService/internal/entity"
	"bannerService/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (h *Handler) getBanner(ctx echo.Context) error {
	err := h.checkUserAuthorization(ctx.Request().Header.Get("Authorization"))
	if err != nil {
		utils.Logger.Error("incorrect auth data", zap.String("error", err.Error()))
		return err
	}
	var banner entity.UserBanner

	if ctx.QueryParam("tag_id") == "" || ctx.QueryParam("feature_id") == "" {
		utils.Logger.Error("incorrect banner data, missing tag or feature")
		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect banner data, missing tag or feature"))
	}

	if err := ctx.Bind(&banner); err != nil {
		utils.Logger.Error("incorrect banner data", zap.String("error", err.Error()))
		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect banner data"))
	}

	result, err := h.services.GetBanner(ctx.Request().Context(), banner)
	if err != nil {
		utils.Logger.Error("get banner error", zap.String("error", err.Error()))
		return responseErr(err)
	}

	return responseOk(ctx, result)
}
