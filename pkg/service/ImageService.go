package service

import "github.com/tumbleweedd/avito-test-task/pkg/repository"

type ImageService struct {
	repository repository.Image
}

func NewImageService(repository repository.Image) *ImageService {
	return &ImageService{
		repository: repository,
	}
}

func (s *ImageService) AddImage(advertisementId int, image string) (int, error) {
	return s.repository.AddImage(advertisementId, image)
}