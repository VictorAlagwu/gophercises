package seed

import (
	"github.com/jinzhu/gorm"
	"github.com/victoralagwu/gophercises/urlshort/mysqlservice/models"
)

var links = []models.Link{
	models.Link{
		Path: "/google",
		Url: "google.com",
	},
	models.Link{
		Path: "/google-api",
		Url: "google.com/search?search=word",
	},
}

func Load(db *gorm.DB) {
	err = db.Debug().DropTableIfExists.(&models.Link{})
}