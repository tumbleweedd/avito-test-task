package service

import (
	advertisement "github.com/tumbleweedd/avito-test-task"
	"github.com/tumbleweedd/avito-test-task/pkg/repository"
)

type ImageService struct {
	repository repository.Image
}

func NewImageService(repository repository.Image) *ImageService {
	return &ImageService{
		repository: repository,
	}
}

func (s *ImageService) GetAllImagesByAdvId(advId int) ([]string, error) {
	return s.repository.GetAllImagesByAdvId(advId)
}

func (s *ImageService) AddImage(advertisementId int, image string) (int, error) {
	return s.repository.AddImage(advertisementId, image)
}

func (s *ImageService) GetImageById(advId, imageId int) (advertisement.ImageResponse, error) {
	return s.repository.GetImageById(advId, imageId)
}

func (s *ImageService) DeleteImage(advId, imageId int) error {
	return s.repository.DeleteImage(advId, imageId)
}
