package repository

import (
	"github.com/render_manga_api/model/entity"
	"log"
)

const (
	PickUp = "pick_up"
	Normal = "normal"
)

type TitleRepository interface {
	GetTitles() (todos []entity.TitleEntity, err error)
}

type titleRepository struct {
}

func NewTitleRepository() TitleRepository {
	return &titleRepository{}
}

func (t *titleRepository) GetTitles() (titles []entity.TitleEntity, err error) {
	titles = []entity.TitleEntity{}
	rows, err := Db.
		Query(`
SELECT id, name, type, thumbnail_url
FROM titles ORDER BY id DESC
`)
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		title := entity.TitleEntity{}
		err = rows.Scan(&title.Id, &title.Name, &title.Type, &title.ThumbnailUrl)
		if err != nil {
			log.Print(err)
			return
		}
		titles = append(titles, title)
	}

	return
}
