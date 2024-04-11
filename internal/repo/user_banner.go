package repo

import (
	"bannerService/internal/utils"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type UserBannerRepo struct {
	db *sqlx.DB
}

func NewUserBanner(db *sqlx.DB) *UserBannerRepo {
	return &UserBannerRepo{db: db}
}

func (r *UserBannerRepo) GetBanner(ctx context.Context, tagId, featureId int) ([]byte, error) {
	bannerQuery := `SELECT content FROM banners WHERE feature_id = $1 AND $2 = ANY (tag_ids) AND is_active`
	row := r.db.QueryRowContext(ctx, bannerQuery, featureId, tagId)

	var content []byte
	err := row.Scan(&content)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(utils.ErrNotFound, fmt.Sprintf("Banner with tag %v and feature %v not found", tagId, featureId))
		}
		return nil, errors.Wrap(err, fmt.Sprintf("UserBannerRepo.GetBanner: %s", err.Error()))
	}

	return content, nil
}
