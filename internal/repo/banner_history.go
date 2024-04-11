package repo

import (
	"bannerService/internal/entity"
	"bannerService/internal/utils"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

const versionsCount = 3

type BannerHistoryRepo struct {
	db *sqlx.DB
}

func NewBannerHistory(db *sqlx.DB) *BannerHistoryRepo {
	return &BannerHistoryRepo{db: db}
}

func (r *BannerHistoryRepo) GetBannerHistory(ctx context.Context, id int) (entity.BannerHistory, error) {
	historyQuery := `SELECT tag_ids, feature_id, content, is_active FROM banners_history
					WHERE banner_id = $1 ORDER BY id DESC LIMIT $2`

	rows, err := r.db.QueryContext(ctx, historyQuery, id, versionsCount)
	if err != nil {
		return entity.BannerHistory{}, errors.Wrap(err, fmt.Sprintf("BannerHistoryRepo.GetBannersHistory: %s", err.Error()))
	}

	banners := entity.BannerHistory{
		BannerId: id,
		Versions: make([]entity.BannerVersion, 0, versionsCount),
	}
	index := versionsCount - 1
	for rows.Next() {
		banner := entity.BannerVersion{
			Version: index + 1,
		}
		content := entity.BannerContent{}

		err = rows.Scan(pq.Array(&banner.TagIds), &banner.FeatureId, &content.Content, &banner.IsActive)
		if err != nil {
			utils.Logger.Error(fmt.Sprintf("BannerHistoryRepo.GetBannersHistory: %s", err.Error()))
			break
		}

		err = json.Unmarshal(content.Content, &banner.Content)
		if err != nil {
			utils.Logger.Error(fmt.Sprintf("BannerHistoryRepo.GetBannersHistory: %s", err.Error()))
			continue
		}

		if index < 0 {
			break
		}
		banners.Versions = append(banners.Versions, banner)
		index--
	}

	if err = rows.Close(); err != nil {
		return entity.BannerHistory{}, err
	}

	return banners, nil
}

func (r *BannerHistoryRepo) SetBannerVersion(ctx context.Context, id, version int) error {
	recentBannerQuery := `SELECT tag_ids, feature_id, content, is_active FROM banners WHERE banner_id = $1`
	row := r.db.QueryRowContext(ctx, recentBannerQuery, id)

	var recentBanner entity.OldBanner
	if err := row.Scan(pq.Array(&recentBanner.TagIds), &recentBanner.FeatureId,
		&recentBanner.Content, &recentBanner.IsActive); err != nil {
		return errors.Wrap(err, fmt.Sprintf("BannerHistoryRepo.SetBannerVersion: %s", err.Error()))
	}

	bannerVersionQuery := `SELECT tag_ids, feature_id, content, is_active FROM banners_history
							WHERE banner_id = $1 ORDER BY id DESC LIMIT 1 OFFSET $2`
	row = r.db.QueryRowContext(ctx, bannerVersionQuery, id, versionsCount-version)

	var bannerFromHistory entity.OldBanner
	if err := row.Scan(pq.Array(&bannerFromHistory.TagIds), &bannerFromHistory.FeatureId,
		&bannerFromHistory.Content, &bannerFromHistory.IsActive); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.Wrap(utils.ErrBadRequest, fmt.Sprintf("no banner %d version %d in history", id, version))
		}
		return errors.Wrap(err, fmt.Sprintf("BannerHistoryRepo.SetBannerVersion: %s", err.Error()))
	}

	if recentBanner.Equals(bannerFromHistory) {
		return errors.Wrap(utils.ErrBadRequest,
			fmt.Sprintf("BannerHistoryRepo.SetBannerVersion: banner version to set must be different from recent banner"))
	}

	tx, err := r.db.Begin()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("BannerHistoryRepo.SetBannerVersion: %s", err.Error()))
	}

	setBannerQuery := `UPDATE banners SET tag_ids = $1, feature_id = $2, content = $3, is_active = $4, updated_at = NOW() 
               		WHERE banner_id = $5`
	_, err = tx.ExecContext(ctx, setBannerQuery, pq.Array(bannerFromHistory.TagIds), bannerFromHistory.FeatureId,
		bannerFromHistory.Content, bannerFromHistory.IsActive, id)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, fmt.Sprintf("BannerHistoryRepo.SetBannerVersion: %s", err.Error()))
	}

	insertOldBanner := `INSERT INTO banners_history (banner_id, content, tag_ids, feature_id, is_active) 
					VALUES ($1, $2, $3, $4, $5)`
	_, err = tx.ExecContext(ctx, insertOldBanner, id, recentBanner.Content, pq.Array(recentBanner.TagIds),
		recentBanner.FeatureId, recentBanner.IsActive)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, fmt.Sprintf("BannerHistoryRepo.SetBannerVersion: %s", err.Error()))
	}

	return tx.Commit()
}
