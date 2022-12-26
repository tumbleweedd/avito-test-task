package service

import (
	"github.com/tumbleweedd/avito-test-task"
	"github.com/tumbleweedd/avito-test-task/pkg/repository"
)

type Advertisement interface {
	CreateAdvertisement(input advertisement.Advertisement) (int, error)
	GetAllAdvertisement() ([]advertisement.Advertisement, error)
	GetAdvertisementById(id int) (advertisement.AdvertisementDTO, error)
}

type Image interface {
	AddImage(advertisementId int, image string) (int, error)
}

type Service struct {
	Advertisement
	Image
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Advertisement: NewAdvertisementService(r.Advertisement),
		Image:         NewImageService(r.Image),
	}
}