package Model

import (
	"time"
)

type Feed struct {
	Id        int    `gorm:"primaryKey" json:"feed_id"`
	Date      string `gorm:"size:255" json:"date"`
	Url       string `json:"url"`
	Link_id   int
	Title     string    `gorm:"size:255" json:"title"`
	UniqueId  string    `gorm:"size:255" json:"unique_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func init() {
	// Db.AutoMigrate(Url{})
}
