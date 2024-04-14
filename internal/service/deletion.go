package service

import (
	"bannerService/internal/deletion"
	"bannerService/internal/repo"
	"context"
	"time"
)

type DeletionService struct {
	repo      repo.Deletion
	queue     *deletion.DeletionQueue
	closeChan chan struct{}
}

func NewDeletionService(repo repo.Deletion, queue *deletion.DeletionQueue) *DeletionService {
	s := &DeletionService{
		repo:      repo,
		queue:     queue,
		closeChan: make(chan struct{}, 1),
	}

	go s.DeletionWorker(context.Background())

	return s
}

func (s *DeletionService) DeletionWorker(ctx context.Context) {
	for {
		select {
		case <-s.closeChan:
			return
		case <-time.After(10 * time.Second):
			s.repo.DeleteFromDB(ctx, s.queue.GetDeletionDataAndClearQueue())
		}
	}
}

func (s *DeletionService) Close() error {
	close(s.closeChan)
	return s.repo.DeleteFromDB(context.Background(), s.queue.GetDeletionDataAndClearQueue())
}

func (s *DeletionService) AddFeatureToDeletionQueue(ctx context.Context, featureId int) {
	s.queue.AddFeatureToQueue(featureId)
}

func (s *DeletionService) AddTagToDeletionQueue(ctx context.Context, tagId int) {
	s.queue.AddTagToQueue(tagId)
}

func (s *DeletionService) AddIdToDeletionQueue(ctx context.Context, id int) {
	s.queue.AddIdToQueue(id)
}
