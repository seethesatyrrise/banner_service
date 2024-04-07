package repo

import (
	"bannerService/internal/entity"
	"bannerService/internal/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"strings"
)

type BannerRepo struct {
	db *sqlx.DB
}

func NewBanner(db *sqlx.DB) *BannerRepo {
	return &BannerRepo{db: db}
}

func (r *BannerRepo) CreateBanner(ctx context.Context, banner entity.Banner) (int, error) {
	content, err := json.Marshal(banner.Content)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("BannerRepo.CreateBanner: %s", err.Error()))
	}

	tx, err := r.db.Begin()
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("BannerRepo.CreateBanner: %s", err.Error()))
	}

	bannerQuery := "INSERT INTO banners (content, is_active) VALUES ($1, $2) RETURNING id"
	row := tx.QueryRowContext(ctx, bannerQuery, content, banner.IsActive)

	var bannerId int
	if err = row.Scan(&bannerId); err != nil {
		tx.Rollback()
		return 0, errors.Wrap(err, fmt.Sprintf("BannerRepo.CreateBanner: %s", err.Error()))
	}

	bannerRelationsQuery := "INSERT INTO feature_tag_banners (tag_id, feature_id, banner_id) VALUES "
	values := make([]interface{}, 0, 3*len(banner.TagIds))

	strBuilder := strings.Builder{}
	strBuilder.WriteString(bannerRelationsQuery)

	for i, tagId := range banner.TagIds {
		strBuilder.WriteString(fmt.Sprintf("($%d, $%d, $%d),", 3*i+1, 3*i+2, 3*i+3))
		values = append(values, tagId, banner.FeatureId, bannerId)
	}
	bannerRelationsQuery = strBuilder.String()[:strBuilder.Len()-1]

	utils.Logger.Info(fmt.Sprintf("BannerRepo.CreateBanner: generated query:\n %s", bannerRelationsQuery))

	res, err := tx.Exec(bannerRelationsQuery, values...)
	if err != nil {
		tx.Rollback()
		return 0, errors.Wrap(err, fmt.Sprintf("BannerRepo.CreateBanner: %s", err.Error()))
	}
	rowsAdded, err := res.RowsAffected()
	utils.Logger.Info(fmt.Sprintf("BannerRepo.CreateBanner: add %d rows", rowsAdded))

	return bannerId, tx.Commit()
}
