package dto

type PagesRequest struct {
	TitleId int `json:"title_id"`
}

type PagesResponse struct {
	ImageUrls []string `json:"page_urls"`
}
