package handler

import (
	"bannerService/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type token struct {
	userToken  string
	adminToken string
}

func (h *Handler) checkUserAuth(fn func(ctx echo.Context) error) func(echo.Context) error {
	return func(ctx echo.Context) error {
		tkn := ctx.Request().Header.Get("Authorization")
		if tkn == "" {
			return responseErr(errors.Wrap(utils.ErrNoAuthorization, "authorization required"))
		}
		if tkn != h.tokens.userToken && tkn != h.tokens.adminToken {
			return responseErr(errors.Wrap(utils.ErrAccessDenied, "access denied"))
		}
		return fn(ctx)
	}
}

func (h *Handler) checkAdminAuth(fn func(ctx echo.Context) error) func(echo.Context) error {
	return func(ctx echo.Context) error {
		tkn := ctx.Request().Header.Get("Authorization")
		if tkn == "" {
			return responseErr(errors.Wrap(utils.ErrNoAuthorization, "authorization required"))
		}
		if tkn != h.tokens.adminToken {
			return responseErr(errors.Wrap(utils.ErrAccessDenied, "access denied"))
		}
		return fn(ctx)
	}
}
