package handler

import (
	"bannerService/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

type ResponseId struct {
	BannerId int `json:"banner_id"`
}

func responseOk(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, data)
}

func responseCreated(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusCreated, data)
}

func responseDeleted(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusNoContent, data)
}

func responseErr(err error) *echo.HTTPError {
	switch errors.Cause(err) {
	case utils.ErrNotFound:
		return echo.NewHTTPError(http.StatusNotFound, err)
	case utils.ErrAlreadyExists:
		return echo.NewHTTPError(http.StatusConflict, err)
	case utils.ErrBadRequest:
		return echo.NewHTTPError(http.StatusBadRequest, err)
	case utils.ErrAccessDenied:
		return echo.NewHTTPError(http.StatusForbidden, err)
	case utils.ErrNoAuthorization:
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
}
