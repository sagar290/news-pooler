package Services

import (
	models "news-poolerr/Model"
	structs "news-poolerr/Structs"
	"time"

	"github.com/gocolly/colly"
)

func FetchNews(link models.Link) {
	var newses []structs.Feed

	collector := colly.NewCollector()

	collector.OnHTML(link.Section, func(element *colly.HTMLElement) {
		// href := element.Attr(link.UrlSelector)

		href := element.ChildAttr(link.UrlSelector, "href")
		// factDesc := element.ChildAttr("h2", "home-title").
		title := element.ChildText(link.DescriptionSelector)
		currentTime := time.Now()
		news := structs.Feed{
			Date:    currentTime.Format("2006-01-02"),
			Link_id: link.Id,
			Url:     href,
			Title:   title,
		}

		newses = append(newses, news)

	})

	collector.Visit(link.Url)

	for _, news := range newses {
		StoreFeed(news)
	}

}

func StoreFeed(data structs.Feed) {
	var feed models.Feed

	models.Db.Where(&models.Feed{
		Date:  data.Date,
		Title: data.Title,
	}).Attrs(models.Feed{
		Date:    data.Date,
		Link_id: data.Link_id,
		Url:     data.Url,
		Title:   data.Title,
	}).FirstOrCreate(&feed)
}


func GetDateList(linkId int) []structs.Date  {
	var dates []structs.Date
	models.Db.Raw("SELECT DISTINCT date FROM feeds WHERE link_id = ? ORDER BY date DESC", linkId).Scan(&dates)
	return dates
}

func GetFeedsByDate(date string) []models.Feed  {
	var feeds []models.Feed

	result := models.Db.Model(&models.Feed{}).Where("date LIKE ?", date).Find(&feeds)

	if result.Error != nil {
		return feeds
	}

	return feeds

}