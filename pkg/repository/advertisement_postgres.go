package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/avito-test-task/model"
)

type AdvertisementPostgres struct {
	db *sqlx.DB
}

func NewAdvertisementPostgres(db *sqlx.DB) *AdvertisementPostgres {
	return &AdvertisementPostgres{
		db: db,
	}
}

func (r *AdvertisementPostgres) DeleteAdvertisement(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	queryForAdv := fmt.Sprint(`
									delete from advertisement a 
       								where a.id = $1
       								`)

	queryForImg := fmt.Sprint(`
									delete
									from img
									where img.id in (select ai.img_id
													 from advertisement_img ai
													 where ai.img_id in (select ai2.img_id from advertisement_img ai2 where ai2.advertisement_id = $1));
									`)

	_, err = tx.Exec(queryForImg, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(queryForAdv, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *AdvertisementPostgres) CreateAdvertisement(input model.Advertisement) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var advertisementId int
	createAdvertisementQuery := fmt.Sprint(`insert into advertisement (title, description, date_creation, price) 
											values($1, $2, current_timestamp, $3) returning id`)

	rowA := tx.QueryRow(createAdvertisementQuery, input.Title, input.Description, input.Price)
	err = rowA.Scan(&advertisementId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var imgId int
	createImgQuery := fmt.Sprint("insert into img (img) values($1) returning id")

	rowA = tx.QueryRow(createImgQuery, input.Img)
	err = rowA.Scan(&imgId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createAdvertisementImgQuery := fmt.Sprint("insert into advertisement_img (advertisement_id, img_id) values ($1,$2)")
	_, err = tx.Exec(createAdvertisementImgQuery, advertisementId, imgId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return advertisementId, tx.Commit()
}

func (r *AdvertisementPostgres) GetAllAdvertisement(sortParam string, limitParam, offsetParam int) ([]model.Advertisement, error) {
	var advertisements []model.Advertisement

	query := fmt.Sprintf(`
						select * from %s as v
						order by %s
						limit %d
						offset %d;
						`, getAllAdvertisementView, sortParam, limitParam, offsetParam)
	if err := r.db.Select(&advertisements, query); err != nil {
		return nil, err
	}

	return advertisements, nil
}

func (r *AdvertisementPostgres) GetAdvertisementById(id int) (model.AdvertisementDTO, error) {
	var advertisement model.AdvertisementDTO

	query := fmt.Sprint(`
								SELECT a.description,
									   a.title,
									   i.img,
									   a.date_creation,
									   a.price
									FROM advertisement a
									join advertisement_img ai on a.id = ai.advertisement_id
									join img i on i.id = ai.img_id
									where a.id = $1
									ORDER  BY i.id, a.id
									limit (select count(*) from advertisement)
									`)

	err := r.db.Get(&advertisement, query, id)

	return advertisement, err

}

func (r *AdvertisementPostgres) UpdateAdvertisement(id int, dto model.UpdateAdvertisement) error {
	queryAdv := fmt.Sprint(`
						update advertisement
						set description = $1,
							title       = $2
						where id = $3
						`)

	queryImg := fmt.Sprint(`
						update img
						set img = $1
						where img.id = (SELECT i.id
											FROM advertisement a
											join advertisement_img ai on a.id = ai.advertisement_id
											join img i on i.id = ai.img_id
											where a.id = $2
											ORDER BY i.id, a.id
											limit 1)
						`)

	_, err := r.db.Exec(queryAdv, dto.Description, dto.Title, id)

	_, err = r.db.Exec(queryImg, dto.Img, id)

	return err

}
