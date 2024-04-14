package handler

import (
	"bannerService/internal/entity"
	"bannerService/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strconv"
)

func (h *Handler) getBannerHistory(ctx echo.Context) error {
	var err error
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

func (h *Handler) setBannerVersion(ctx echo.Context) error {
	var id entity.BannerId
	var version int
	var err error

	if err := ctx.Bind(&id); err != nil {
		utils.Logger.Error("incorrect data", zap.String("error", err.Error()))
		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect data"))
	}

	if v := ctx.QueryParam("set_version"); v != "" {
		version, err = strconv.Atoi(v)
	}

	if err != nil || version < 1 || version > 3 {
		utils.Logger.Error("incorrect data: version must be from 1 to 3")
		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect data: version must be from 1 to 3"))
	}

	err = h.services.SetBannerVersion(ctx.Request().Context(), id.BannerId, version)
	if err != nil {
		utils.Logger.Error("set banner version error", zap.String("error", err.Error()))
		return responseErr(err)
	}

	return responseOk(ctx, ResponseMessage{"ok"})
}
