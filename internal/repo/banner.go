package repo

import (
	"bannerService/internal/entity"
	"bannerService/internal/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
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

	bannerQuery := `INSERT INTO banners (content, tag_ids, feature_id, is_active) 
					VALUES ($1, $2, $3, $4)
					RETURNING banner_id`
	row := r.db.QueryRowContext(ctx, bannerQuery, content, pq.Array(banner.TagIds), banner.FeatureId, banner.IsActive)

	var bannerId int
	if err = row.Scan(&bannerId); err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("BannerRepo.CreateBanner: %s", err.Error()))
	}

	return bannerId, nil
}

func (r *BannerRepo) DeleteBanner(ctx context.Context, id int) error {
	deleteQuery := `DELETE FROM banners WHERE banner_id = $1;`

	res, err := r.db.ExecContext(ctx, deleteQuery, id)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("BannerRepo.DeleteBanner: %s", err.Error()))
	}

	rowsDeleted, _ := res.RowsAffected()
	if rowsDeleted == 0 {
		return errors.Wrap(utils.ErrNotFound, fmt.Sprintf("BannerRepo.DeleteBanner: no banners for id %d", id))
	}
	utils.Logger.Info(fmt.Sprintf("BannerRepo.DeleteBanner: delete %d rows from banners table", rowsDeleted))

	return nil
}

func (r *BannerRepo) FilterBanners(ctx context.Context, params entity.BannerFilters) ([]entity.BannerInfo, error) {

	filterQuery := `SELECT *  FROM banners
					WHERE ($1 = 0 OR $1 = ANY (tag_ids)) AND ($2 = 0 OR $2 = feature_id)
					LIMIT (CASE WHEN $3 > 0 THEN $3 END) OFFSET $4`

	rows, err := r.db.QueryContext(ctx, filterQuery, params.TagId, params.FeatureId, params.Limit, params.Offset)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("BannerRepo.FilterBanners: %s", err.Error()))
	}

	banners := []entity.BannerInfo{}
	for rows.Next() {
		banner := entity.BannerInfo{}
		content := entity.BannerContent{}

		err = rows.Scan(&banner.BannerId, &content.Content, pq.Array(&banner.TagIds), &banner.FeatureId, &banner.IsActive,
			&banner.CreatedAt, &banner.UpdatedAt)
		if err != nil {
			utils.Logger.Error(fmt.Sprintf("BannerRepo.FilterBanners: %s", err.Error()))
			break
		}

		err = json.Unmarshal(content.Content, &banner.Content)
		if err != nil {
			utils.Logger.Error(fmt.Sprintf("BannerRepo.FilterBanners: %s", err.Error()))
			continue
		}

		banners = append(banners, banner)
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return banners, nil
}

func (r *BannerRepo) UpdateBanner(ctx context.Context, patch map[string]interface{}) error {
	bannerId, ok := patch["banner_id"]
	if !ok {
		return errors.New(fmt.Sprintf("BannerRepo.UpdateBanner: no id"))
	}

	oldBannerQuery := `SELECT tag_ids, feature_id, content, is_active FROM banners WHERE banner_id = $1`
	row := r.db.QueryRowContext(ctx, oldBannerQuery, bannerId)

	var oldBanner entity.OldBanner
	if err := row.Scan(pq.Array(&oldBanner.TagIds), &oldBanner.FeatureId, &oldBanner.Content, &oldBanner.IsActive); err != nil {
		return errors.Wrap(err, fmt.Sprintf("BannerRepo.UpdateBanner: %s", err.Error()))
	}

	tx, err := r.db.Begin()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("BannerRepo.UpdateBanner: %s", err.Error()))
	}

	historyQuery := `INSERT INTO banners_history (banner_id, content, tag_ids, feature_id, is_active) 
					VALUES ($1, $2, $3, $4, $5)`
	_, err = tx.ExecContext(ctx, historyQuery, bannerId, oldBanner.Content, pq.Array(oldBanner.TagIds),
		oldBanner.FeatureId, oldBanner.IsActive)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, fmt.Sprintf("BannerRepo.UpdateBanner: %s", err.Error()))
	}

	if content, ok := patch["content"]; ok {
		contentJSON, err := json.Marshal(content)
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, fmt.Sprintf("BannerRepo.UpdateBanner: %s", err.Error()))
		}
		oldBanner.Content = contentJSON
		delete(patch, "content")
	}
	patchJSON, err := json.Marshal(patch)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, fmt.Sprintf("BannerRepo.UpdateBanner: %s", err.Error()))
	}

	err = json.Unmarshal(patchJSON, &oldBanner)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, fmt.Sprintf("BannerRepo.UpdateBanner: %s", err.Error()))
	}

	bannerQuery := `UPDATE banners SET tag_ids = $1, feature_id = $2, content = $3, is_active = $4, updated_at = NOW() 
               		WHERE banner_id = $5`
	_, err = tx.ExecContext(ctx, bannerQuery, pq.Array(oldBanner.TagIds), oldBanner.FeatureId, oldBanner.Content, oldBanner.IsActive, bannerId)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, fmt.Sprintf("BannerRepo.UpdateBanner: %s", err.Error()))
	}

	return tx.Commit()
}
