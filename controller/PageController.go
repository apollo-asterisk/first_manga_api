package controller

import (
	"encoding/json"
	"github.com/render_manga_api/controller/dto"
	"github.com/render_manga_api/model/repository"
	"log"
	"net/http"
	"strconv"
)

type PageController interface {
	GetPages(w http.ResponseWriter, r *http.Request)
}

type pageController struct {
	pr repository.PageRepository
}

func NewPageController(pr repository.PageRepository) PageController {
	return &pageController{pr}
}

func (c pageController) GetPages(w http.ResponseWriter, r *http.Request) {
	titleId, err := strconv.Atoi(r.URL.Query().Get("title_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	pageImageUrls, err := c.pr.GetPageImageUrlsByTitleId(titleId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var pagesResponse dto.PagesResponse
	pagesResponse = dto.PagesResponse{ImageUrls: pageImageUrls}

	output, err := json.MarshalIndent(pagesResponse, "", "\t\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

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
