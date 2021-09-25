package Model

import (
	"time"
)

type Link struct {
	Id                  int       `gorm:"primaryKey" json:"feed_id"`
	Title               string    `gorm:"size:255" json:"title"`
	Thumbnail           string    `json:"thumbnail"`
	Url                 string    `json:"url"`
	UrlSelector         string    `gorm:"size:255" json:"url_selector"`
	DescriptionSelector string    `gorm:"size:255" json:"description_selector"`
	UniqueId            string    `gorm:"size:255" json:"unique_id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func init() {
	// Db.AutoMigrate(Click{})
}
