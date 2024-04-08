package handler

import (
	"bannerService/internal/utils"
	"github.com/pkg/errors"
)

type token struct {
	userToken  string
	adminToken string
}

func (h *Handler) checkUserAuthorization(token string) error {
	if token == "" {
		return responseErr(errors.Wrap(utils.ErrNoAuthorization, "authorization required"))
	}
	if token != h.tokens.userToken && token != h.tokens.adminToken {
		return responseErr(errors.Wrap(utils.ErrAccessDenied, "access denied"))
	}
	return nil
}

func (h *Handler) checkAdminAuthorization(token string) error {
	if token == "" {
		return responseErr(errors.Wrap(utils.ErrNoAuthorization, "authorization required"))
	}
	if token != h.tokens.adminToken {
		return responseErr(errors.Wrap(utils.ErrAccessDenied, "access denied"))
	}
	return nil
}
