package service

import (
	"github.com/tumbleweedd/avito-test-task/model"
	"github.com/tumbleweedd/avito-test-task/pkg/repository"
)

type Advertisement interface {
	CreateAdvertisement(input model.Advertisement) (int, error)
	GetAllAdvertisement(sortParam string, limitParam, offsetParam int) ([]model.Advertisement, error)
	GetAdvertisementById(id int) (model.AdvertisementDTO, error)
	UpdateAdvertisement(id int, dto model.UpdateAdvertisement) error
	DeleteAdvertisement(id int) error
}

type Image interface {
	AddImage(advertisementId int, image string) (int, error)
	GetAllImagesByAdvId(advId int) ([]string, error)
	GetImageById(advId, imageId int) (model.ImageResponse, error)
	DeleteImage(advId, imageId int) error
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
