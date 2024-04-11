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

type BannerHistoryRepo struct {
	db *sqlx.DB
}

func NewBannerHistory(db *sqlx.DB) *BannerHistoryRepo {
	return &BannerHistoryRepo{db: db}
}

func (r *BannerHistoryRepo) GetBannerHistory(ctx context.Context, id int) (entity.BannerHistory, error) {
	historyQuery := `SELECT tag_ids, feature_id, content, is_active FROM banners_history
					WHERE banner_id = $1 ORDER BY id DESC LIMIT $2`

	rows, err := r.db.QueryContext(ctx, historyQuery, id, entity.VersionsCount)
	if err != nil {
		return entity.BannerHistory{}, errors.Wrap(err, fmt.Sprintf("BannerHistoryRepo.GetBannersHistory: %s", err.Error()))
	}

	banners := entity.BannerHistory{
		BannerId: id,
	}
	index := entity.VersionsCount - 1
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
		banners.Versions[index] = banner
		index--
	}

	if err = rows.Close(); err != nil {
		return entity.BannerHistory{}, err
	}

	return banners, nil
}
