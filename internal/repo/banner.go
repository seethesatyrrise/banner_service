package repo

import (
	"bannerService/internal/entity"
	"bannerService/internal/utils"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type BannerRepo struct {
	db *sqlx.DB
}

func NewBanner(db *sqlx.DB) *BannerRepo {
	return &BannerRepo{db: db}
}

func (r *BannerRepo) AddBanner(ctx context.Context, banner entity.BannerToDB) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("BannerRepo.AddBanner: %s", err.Error()))
	}

	bannerQuery := "INSERT INTO banners (content, is_active) VALUES ($1, $2) RETURNING id"
	row := tx.QueryRowContext(ctx, bannerQuery, banner.Content, banner.IsActive)

	var bannerId int
	if err = row.Scan(&bannerId); err != nil {
		tx.Rollback()
		return 0, errors.Wrap(err, fmt.Sprintf("BannerRepo.AddBanner: %s", err.Error()))
	}

	return bannerId, tx.Commit()
}

func (r *BannerRepo) AddBannerRelations(ctx context.Context, bannerRelations []entity.BannerRelationsToDB) error {
	tx, err := r.db.Begin()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("BannerRepo.AddBannerRelations: %s", err.Error()))
	}

	bannerRelationsQuery := "INSERT INTO feature_tag_banners (tag_id, feature_id, banner_id) VALUES (:tag_id, :feature_id, :banner_id)"

	res, err := r.db.NamedExecContext(ctx, bannerRelationsQuery, bannerRelations)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, fmt.Sprintf("BannerRepo.AddBannerRelations: %s", err.Error()))
	}

	rowsAdded, _ := res.RowsAffected()
	utils.Logger.Info(fmt.Sprintf("BannerRepo.AddBannerRelations: add %d rows", rowsAdded))

	return tx.Commit()
}
