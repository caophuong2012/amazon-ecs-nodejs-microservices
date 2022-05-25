package request

import "mime/multipart"

type CreateStoreFront struct {
	CreatorID                     int64          `json:"creator_id"`
	StoreName                     string         `json:"store_name"`
	StoreUrl                      string         `json:"store_url"`
	IsUrl                         string         `json:"is_url"`
	StoreFrontTitleHeader         string         `json:"storefront_title_header"`
	StoreFrontSubHeader           string         `json:"storefront_sub_header"`
	PromotionalMessage            string         `json:"promotional_message"`
	StoreLogoFile                 multipart.File `json:"store_logo_file"`
	StoreLogoFileName             string         `json:"store_logo_filename"`
	StoreFrontBannerVideoFile     multipart.File `json:"storefront_banner_video_file"`
	StoreFrontBannerVideoFileName string         `json:"storefront_banner_video_filename"`
	StoreFrontBannerVideoUrl      string         `json:"storefront_banner_video_url"`
}
