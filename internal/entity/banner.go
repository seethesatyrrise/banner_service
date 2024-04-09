package entity

import "time"

type Banner struct {
	TagIds    []int                  `json:"tag_ids" db:"tag_ids"`
	FeatureId int                    `json:"feature_id" db:"feature_id"`
	Content   map[string]interface{} `json:"content" db:"content"`
	IsActive  bool                   `json:"is_active" db:"is_active"`
}

type BannerToDB struct {
	Content  []byte
	IsActive bool
}

type BannerRelationsToDB struct {
	TagId     int `db:"tag_id"`
	FeatureId int `db:"feature_id"`
	BannerId  int `db:"banner_id"`
}

type UserBanner struct {
	TagId           int  `query:"tag_id"`
	FeatureId       int  `query:"feature_id"`
	UseLastRevision bool `query:"use_last_revision"`
}

type BannerId struct {
	Id int `param:"id"`
}

type BannerInfo struct {
	BannerId  int                    `json:"banner_id"`
	TagIds    []int                  `json:"tag_ids"`
	FeatureId int                    `json:"feature_id"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}
