package service

import (
	"github.com/tumbleweedd/avito-test-task"
	"github.com/tumbleweedd/avito-test-task/pkg/repository"
)

type AdvertisementService struct {
	repo repository.Advertisement
}

func (s *AdvertisementService) DeleteAdvertisement(id int) error {
	return s.repo.DeleteAdvertisement(id)
}

func NewAdvertisementService(repo repository.Advertisement) *AdvertisementService {
	return &AdvertisementService{
		repo: repo,
	}
}

func (s *AdvertisementService) CreateAdvertisement(input advertisement.Advertisement) (int, error) {
	return s.repo.CreateAdvertisement(input)
}

func (s *AdvertisementService) GetAllAdvertisement() ([]advertisement.Advertisement, error) {
	return s.repo.GetAllAdvertisement()
}

func (s *AdvertisementService) GetAdvertisementById(id int) (advertisement.AdvertisementDTO, error) {
	return s.repo.GetAdvertisementById(id)
}

func (s *AdvertisementService) UpdateAdvertisement(id int, dto advertisement.UpdateAdvertisement) error {
	if err := dto.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateAdvertisement(id, dto)
}
