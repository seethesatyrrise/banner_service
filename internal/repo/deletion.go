package repo

import (
	"bannerService/internal/entity"
	"bannerService/internal/utils"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type DeletionRepo struct {
	db *sqlx.DB
}

func NewDeletion(db *sqlx.DB) *DeletionRepo {
	return &DeletionRepo{db: db}
}

func (r *DeletionRepo) DeleteFromDB(ctx context.Context, data entity.Deletion) error {
	deleteQuery := `DELETE FROM banners WHERE banner_id = ANY($1) OR feature_id = ANY($2) OR (tag_ids && $3);`

	res, err := r.db.ExecContext(ctx, deleteQuery, pq.Array(data.Ids), pq.Array(data.Features), pq.Array(data.Tags))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("DeletionRepo.DeleteFromDB: %s", err.Error()))
	}

	rowsDeleted, _ := res.RowsAffected()
	utils.Logger.Info(fmt.Sprintf("DeletionRepo.DeleteFromDB: delete %d rows from banners table", rowsDeleted))

	return nil
}
