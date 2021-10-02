package Model

import (
	"time"
)

type Link struct {
	Id                  int       `gorm:"primaryKey" json:"link_id"`
	Title               string    `gorm:"size:255" json:"title"`
	Thumbnail           string    `json:"thumbnail"`
	Url                 string    `json:"url"`
	Section             string    `gorm:"size:255" json:"section"`
	UrlSelector         string    `gorm:"size:255" json:"url_selector"`
	DescriptionSelector string    `gorm:"size:255" json:"description_selector"`
	UniqueId            string    `gorm:"size:255" json:"unique_id"`
	Feeds               []Feed    `gorm:"foreignKey:Link_id;references:Id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func init() {
	// Db.AutoMigrate(Click{})
}
