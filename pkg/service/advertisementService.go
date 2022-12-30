package service

import (
	"github.com/tumbleweedd/avito-test-task/model"
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

func (s *AdvertisementService) CreateAdvertisement(input model.Advertisement) (int, error) {
	return s.repo.CreateAdvertisement(input)
}

func (s *AdvertisementService) GetAllAdvertisement(sortParam string) ([]model.Advertisement, error) {
	return s.repo.GetAllAdvertisement(sortParam)
}

func (s *AdvertisementService) GetAdvertisementById(id int) (model.AdvertisementDTO, error) {
	return s.repo.GetAdvertisementById(id)
}

func (s *AdvertisementService) UpdateAdvertisement(id int, dto model.UpdateAdvertisement) error {
	if err := dto.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateAdvertisement(id, dto)
}
