package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/avito-test-task/model"
)

type Advertisement interface {
	GetAllAdvertisement() ([]model.Advertisement, error)
	CreateAdvertisement(input model.Advertisement) (int, error)
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

type Repository struct {
	Advertisement
	Image
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Advertisement: NewAdvertisementPostgres(db),
		Image:         NewImagePostgres(db),
	}
}
