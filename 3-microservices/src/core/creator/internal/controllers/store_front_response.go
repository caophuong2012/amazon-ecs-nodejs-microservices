package controllers

type StoreFrontResponse struct {
	CreatorID                     int64  `json:"creator_id"`
	StoreName                     string `json:"store_name"`
	StoreUrl                      string `json:"store_url"`
	StoreFrontTitleHeader         string `json:"store_front_title_header"`
	StoreFrontSubHeader           string `json:"store_front_sub_header"`
	PromotionalMessage            string `json:"promotional_message"`
	StoreLogoFileName             string `json:"store_logo_file_name"`
	StoreFrontBannerVideoFileName string `json:"store_front_banner_video_file_name"`
}
