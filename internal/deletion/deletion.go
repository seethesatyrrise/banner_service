package deletion

import (
	"bannerService/internal/entity"
	"sync"
)

type DeletionQueue struct {
	Features entity.FeaturesQueue
	Tags     entity.TagsQueue
	Ids      entity.IdsQueue
}

func CreateQueue() *DeletionQueue {
	return &DeletionQueue{
		Features: entity.FeaturesQueue{
			Features: make([]int, 0, 10),
			Mtx:      sync.Mutex{},
		},
		Tags: entity.TagsQueue{
			Tags: make([]int, 0, 10),
			Mtx:  sync.Mutex{},
		},
		Ids: entity.IdsQueue{
			Ids: make([]int, 0, 10),
			Mtx: sync.Mutex{},
		},
	}
}

func (q *DeletionQueue) AddFeatureToQueue(featureId int) {
	q.Features.Mtx.Lock()
	q.Features.Features = append(q.Features.Features, featureId)
	q.Features.Mtx.Unlock()
}

func (q *DeletionQueue) AddTagToQueue(tagId int) {
	q.Tags.Mtx.Lock()
	q.Tags.Tags = append(q.Tags.Tags, tagId)
	q.Tags.Mtx.Unlock()
}

func (q *DeletionQueue) AddIdToQueue(id int) {
	q.Ids.Mtx.Lock()
	q.Ids.Ids = append(q.Ids.Ids, id)
	q.Ids.Mtx.Unlock()
}

func (q *DeletionQueue) GetDeletionDataAndClearQueue() entity.Deletion {
	var deletion entity.Deletion

	q.Features.Mtx.Lock()
	deletion.Features = q.Features.Features
	q.Features.Features = make([]int, 0, 10)
	q.Features.Mtx.Unlock()

	q.Tags.Mtx.Lock()
	deletion.Tags = q.Tags.Tags
	q.Tags.Tags = make([]int, 0, 10)
	q.Tags.Mtx.Unlock()

	q.Ids.Mtx.Lock()
	deletion.Ids = q.Ids.Ids
	q.Ids.Ids = make([]int, 0, 10)
	q.Ids.Mtx.Unlock()

	return deletion
}
