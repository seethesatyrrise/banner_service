package service

import (
	"bannerService/internal/deletion"
	"bannerService/internal/repo"
	"bannerService/internal/utils"
	"context"
	"time"
)

type DeletionService struct {
	repo               repo.Deletion
	queue              *deletion.DeletionQueue
	deletionWorkerQuit chan struct{}
}

func NewDeletionService(repo repo.Deletion, queue *deletion.DeletionQueue, deletionWorkerChan chan struct{}) *DeletionService {
	s := &DeletionService{
		repo:  repo,
		queue: queue,
	}

	go s.DeletionWorker(context.Background(), deletionWorkerChan)

	return s
}

func (s *DeletionService) DeletionWorker(ctx context.Context, deletionWorkerChan chan struct{}) {
	for {
		select {
		case <-deletionWorkerChan:
			s.repo.DeleteFromDB(ctx, s.queue.GetDeletionDataAndClearQueue())
			break
		case <-time.After(10 * time.Second):
			s.repo.DeleteFromDB(ctx, s.queue.GetDeletionDataAndClearQueue())
		}
	}
	utils.Logger.Info("quitting deletion worker")
	deletionWorkerChan <- struct{}{}
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
