package entity

type Banner struct {
	TagIds    []int                  `json:"tag_ids"`
	FeatureId int                    `json:"feature_id"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
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
