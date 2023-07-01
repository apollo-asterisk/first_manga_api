package repository

import (
	"log"
)

type PageRepository interface {
	GetPageImageUrlsByTitleId(titleId int) (pageImages []string, err error)
}

type pageRepository struct {
}

func NewPageRepository() PageRepository {
	return &pageRepository{}
}

func (c pageRepository) GetPageImageUrlsByTitleId(titleId int) (pageImageUrls []string, err error) {
	rows, err := Db.
		Query(`
SELECT page_image_url
FROM pages
WHERE title_id = $1
ORDER BY id
`, titleId)
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		var pageImageUrl string
		err = rows.Scan(&pageImageUrl)
		if err != nil {
			log.Print(err)
			return
		}
		pageImageUrls = append(pageImageUrls, pageImageUrl)
	}

	return
}
