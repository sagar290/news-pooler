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

func GetLink(link_id int) models.Link {
	var link models.Link

	result := models.Db.Where(&models.Link{Id: link_id}).First(&link)

	if result.Error != nil {
		return link
	}

	return link

}

func UpdateLink(url_id int, body structs.LinkBody) (bool, string) {
	// var url models.Url
	var message string

	result := models.Db.Where(&models.Link{Id: url_id}).Updates(models.Link{
		Title:               body.Title,
		Url:                 body.Url,
		UrlSelector:         body.UrlSelector,
		DescriptionSelector: body.DescriptionSelector,
	})

	if result.Error != nil {
		message = "cant update"
		return false, message
	}

	message = "updated"

	return true, message

}

func DeleteLink(id int) (bool, string) {
	var link models.Link
	var message string

	result := models.Db.Where(&models.Link{Id: id}).Delete(&link)
	if result.Error != nil {
		message = "cant delete"
		return false, message
	}

	message = "deleted"

	return true, message

}
