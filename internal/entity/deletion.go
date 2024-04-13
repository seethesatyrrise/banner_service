package entity

import "sync"

type Feature struct {
	FeatureId int `param:"feature_id"`
}

type FeaturesQueue struct {
	Features []int
	Mtx      sync.Mutex
}

type Tag struct {
	TagId int `param:"tag_id"`
}

type TagsQueue struct {
	Tags []int
	Mtx  sync.Mutex
}
type IdsQueue struct {
	Ids []int
	Mtx sync.Mutex
}

type Deletion struct {
	Features []int `db:"features"`
	Tags     []int `db:"tags"`
	Ids      []int `db:"ids"`
}
