package entity

import "time"

type Banner struct {
	TagIds    []int64                `json:"tag_ids" db:"tag_ids"`
	FeatureId int                    `json:"feature_id" db:"feature_id"`
	Content   map[string]interface{} `json:"content" db:"content"`
	IsActive  bool                   `json:"is_active" db:"is_active"`
}

type UserBanner struct {
	TagId           int  `query:"tag_id"`
	FeatureId       int  `query:"feature_id"`
	UseLastRevision bool `query:"use_last_revision"`
}

type BannerId struct {
	BannerId int `json:"banner_id" param:"banner_id"`
}

type BannerFilters struct {
	TagId     int `query:"tag_id" db:"tag_id"`
	FeatureId int `query:"feature_id" db:"feature_id"`
	Limit     int `query:"limit" db:"limit"`
	Offset    int `query:"offset" db:"offset"`
}

type BannerInfo struct {
	BannerId  int                    `json:"banner_id" db:"banner_id"`
	TagIds    []int64                `json:"tag_ids" db:"tag_ids"`
	FeatureId int                    `json:"feature_id" db:"feature_id"`
	Content   map[string]interface{} `json:"content" db:"content"`
	IsActive  bool                   `json:"is_active" db:"is_active"`
	CreatedAt time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt time.Time              `json:"updated_at" db:"updated_at"`
}

type BannerContent struct {
	Content []byte `db:"content"`
}

type OldBanner struct {
	TagIds    []int64 `json:"tag_ids" db:"tag_ids"`
	FeatureId int     `json:"feature_id" db:"feature_id"`
	Content   []byte  `json:"content" db:"content"`
	IsActive  bool    `json:"is_active" db:"is_active"`
}

type BannerVersion struct {
	Version   int                    `json:"version"`
	TagIds    []int64                `json:"tag_ids" db:"tag_ids"`
	FeatureId int                    `json:"feature_id" db:"feature_id"`
	Content   map[string]interface{} `json:"content" db:"content"`
	IsActive  bool                   `json:"is_active" db:"is_active"`
}

const VersionsCount = 3

type BannerHistory struct {
	BannerId int                          `json:"banner_id" param:"banner_id"`
	Versions [VersionsCount]BannerVersion `json:"versions"`
}
