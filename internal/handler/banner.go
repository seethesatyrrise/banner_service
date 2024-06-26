package handler

import (
	"bannerService/internal/entity"
	"bannerService/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (h *Handler) createBanner(ctx echo.Context) error {
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

func (h *Handler) filterBanners(ctx echo.Context) error {
	var queryParams entity.BannerFilters

	if err := ctx.Bind(&queryParams); err != nil {
		utils.Logger.Error("incorrect data", zap.String("error", err.Error()))
		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect data"))
	}

	bannersInfo, err := h.services.FilterBanners(ctx.Request().Context(), queryParams)
	if err != nil {
		utils.Logger.Error("banner filtration error", zap.String("error", err.Error()))
		return responseErr(err)
	}

	return responseOk(ctx, bannersInfo)
}

func (h *Handler) updateBanner(ctx echo.Context) error {
	var err error
	bannerPatch := make(map[string]interface{}, 0)

	if err := ctx.Bind(&bannerPatch); err != nil {
		utils.Logger.Error("incorrect data", zap.String("error", err.Error()))
		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect data"))
	}

	if len(bannerPatch) <= 1 {
		utils.Logger.Error("incorrect data: nothing to update")
		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect data: nothing to update"))
	}

	err = h.services.UpdateBanner(ctx.Request().Context(), bannerPatch)
	if err != nil {
		utils.Logger.Error("banner filtration error", zap.String("error", err.Error()))
		return responseErr(err)
	}

	return responseOk(ctx, ResponseMessage{"ok"})
}
