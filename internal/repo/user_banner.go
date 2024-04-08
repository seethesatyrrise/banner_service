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
	bannerQuery := `SELECT b.content FROM banners as b 
    			INNER JOIN banners_relations as br ON b.id = br.banner_id 
                 WHERE br.tag_id = $1 AND br.feature_id = $2 AND b.is_active 
                 ORDER BY br.updated_at DESC LIMIT 1`
	row := r.db.QueryRowContext(ctx, bannerQuery, tagId, featureId)

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
