package Services

import (
	models "news-poolerr/Model"
	structs "news-poolerr/Structs"
)

func StoreUrl(link structs.LinkBody) (bool, string) {

	// var response models.Url

	var message string

	dbc := models.Db.Create(&models.Link{
		Title:               link.Title,
		Url:                 link.Url,
		UrlSelector:         link.UrlSelector,
		DescriptionSelector: link.DescriptionSelector,
	})

	if dbc.Error != nil {
		message = "cant create user"
		return false, message
	}

	message = "created"

	return true, message

}

func GetLinks() []models.Link {
	var links []models.Link

	result := models.Db.Where(&models.Link{}).Find(&links)

	if result.Error != nil {
		return links
	}

	return links

}
