package Structs

type Feed struct {
	Date     string `json:"date"`
	Link_id  int    `json:"link_id"`
	Url      string `json:"url"`
	Title    string `json:"title"`
	UniqueId string `json:"unique_id"`
}

type Date struct {
	Date string `json:"date"`
}