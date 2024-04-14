package handler

import (
	"bannerService/internal/entity"
	"bannerService/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (h *Handler) deleteId(ctx echo.Context) error {
	var bannerId entity.BannerId

	if err := ctx.Bind(&bannerId); err != nil {
		utils.Logger.Error("incorrect data", zap.String("error", err.Error()))
		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect data"))
	}

	h.services.AddIdToDeletionQueue(ctx.Request().Context(), bannerId.BannerId)

	return responseOk(ctx, ResponseMessage{"id was added to deletion queue"})
}

func (h *Handler) deleteFeature(ctx echo.Context) error {
	var feature entity.Feature

	if err := ctx.Bind(&feature); err != nil {
		utils.Logger.Error("invalid feature id", zap.String("error", err.Error()))
		return responseErr(errors.Wrap(utils.ErrBadRequest, "invalid feature id"))
	}

	if feature.FeatureId < 0 {
		utils.Logger.Error("invalid feature id")
		return responseErr(errors.Wrap(utils.ErrBadRequest, "invalid feature id"))
	}

	h.services.AddFeatureToDeletionQueue(ctx.Request().Context(), feature.FeatureId)

	return responseOk(ctx, ResponseMessage{"feature was added to deletion queue"})
}

func (h *Handler) deleteTag(ctx echo.Context) error {
	var tag entity.Tag

	if err := ctx.Bind(&tag); err != nil {
		utils.Logger.Error("invalid tag id", zap.String("error", err.Error()))
		return responseErr(errors.Wrap(utils.ErrBadRequest, "invalid tag id"))
	}

	if tag.TagId < 0 {
		utils.Logger.Error("invalid feature id")
		return responseErr(errors.Wrap(utils.ErrBadRequest, "invalid tag id"))
	}

	h.services.AddTagToDeletionQueue(ctx.Request().Context(), tag.TagId)

	return responseOk(ctx, ResponseMessage{"tag was added to deletion queue"})
}
