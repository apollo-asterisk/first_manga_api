package controller

import (
	"encoding/json"
	"github.com/render_manga_api/controller/dto"
	"github.com/render_manga_api/model/repository"
	"log"
	"net/http"
)

type TitleController interface {
	GetTitles(w http.ResponseWriter, r *http.Request)
}

type titleController struct {
	tr repository.TitleRepository
}

func NewTitleController(tr repository.TitleRepository) TitleController {
	return &titleController{tr}
}

func (tc *titleController) GetTitles(w http.ResponseWriter, r *http.Request) {
	titles, err := tc.tr.GetTitles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var pickUpTitleResponses []dto.TitleResponse
	var normalTitleResponses []dto.TitleResponse
	for _, v := range titles {
		if v.Type == repository.PickUp {
			pickUpTitleResponses = append(pickUpTitleResponses, dto.TitleResponse{TitleId: v.Id, Name: v.Name, ThumbnailUrl: v.ThumbnailUrl})
		} else {
			normalTitleResponses = append(normalTitleResponses, dto.TitleResponse{TitleId: v.Id, Name: v.Name, ThumbnailUrl: v.ThumbnailUrl})
		}
	}

	var titlesResponse dto.TitlesResponse
	titlesResponse.PickupTitles = pickUpTitleResponses
	titlesResponse.NormalTitles = normalTitleResponses

	output, _ := json.MarshalIndent(titlesResponse, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	_, err = w.Write(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
