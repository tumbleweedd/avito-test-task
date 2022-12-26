package advertisement

type Advertisement struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Img         string `json:"img"`
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
