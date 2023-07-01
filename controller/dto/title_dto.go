package dto

type TitleRequest struct {
	Name string `json:"name"`
}

type TitleResponse struct {
	TitleId      int    `json:"title_id"`
	Name         string `json:"name"`
	ThumbnailUrl string `json:"image_url"`
}

type TitlesResponse struct {
	PickupTitles []TitleResponse `json:"pick_up_titles"`
	NormalTitles []TitleResponse `json:"normal_titles"`
}
