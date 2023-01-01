package model

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Advertisement struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Img         string    `json:"img"`
	DateTime    time.Time `json:"date_time" db:"date_creation"`
	Price       float64   `json:"price" db:"price"`
}

func (adv *Advertisement) Validate() error {
	return validation.ValidateStruct(
		adv,
		validation.Field(&adv.Description, validation.Length(0, 1000).Error("the description length must be no more than"), validation.Required),
		validation.Field(&adv.Title, validation.Length(0, 200).Error("the title length must be no more than 10")),
		validation.Field(&adv.Price, validation.Required),
	)
}

type AdvertisementDTO struct {
	Id          int       `json:"-"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Img         string    `json:"img"`
	DateTime    time.Time `json:"date_time" db:"date_creation"`
	Price       float64   `json:"price" db:"price"`
}

type AdvertisemenWithAllImgtDTO struct {
	Id          int       `json:"-"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Img         []string  `json:"img"`
	DateTime    time.Time `json:"date_time" db:"date_creation"`
	Price       float64   `json:"price" db:"price"`
}

type Img struct {
	Id  int    `json:"id" db:"id"`
	Img string `json:"img" db:"img"`
}

type AdvertisementImg struct {
	Id              int
	AdvertisementId int
	ImgId           int
}

type UpdateAdvertisement struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Img         *string `json:"img"`
}

func (ua *UpdateAdvertisement) Validate() error {
	if ua.Title == nil && ua.Description == nil && ua.Img == nil {
		return errors.New("update structure has no validate")
	}
	return nil
}

type ImageResponse struct {
	Image string
}
