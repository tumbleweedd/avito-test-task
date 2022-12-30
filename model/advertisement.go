package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Advertisement struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Img         string `json:"img"`
}

func (adv *Advertisement) Validate() error {
	return validation.ValidateStruct(
		adv,
		validation.Field(&adv.Description, validation.Length(0, 1000).Error("the description length must be no more than")),
		validation.Field(&adv.Title, validation.Length(0, 200).Error("the title length must be no more than 10")),
	)
}

type AdvertisementDTO struct {
	Id          int    `json:"-"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Img         string `json:"img"`
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
