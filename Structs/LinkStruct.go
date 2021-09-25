package Structs

type LinkBody struct {
	Title               string `json:"title"`
	Url                 string `json:"url"`
	UrlSelector         string `json:"url_selector"`
	DescriptionSelector string `json:"description_selector"`
	UniqueId            string `json:"unique_id"`
}
