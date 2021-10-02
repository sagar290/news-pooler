package Structs

type LinkBody struct {
	Title               string `json:"title" binding:"required"`
	Url                 string `json:"url" binding:"required"`
	Section             string `json:"section" binding:"required"`
	UrlSelector         string `json:"url_selector" binding:"required"`
	DescriptionSelector string `json:"description_selector" binding:"required"`
	UniqueId            string `json:"unique_id"`
}
