package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/victoralagwu/gophercises/urlshort/mysqlservice/models"
)

var links = []models.Link{
	models.Link{
		Path: "/google",
		Url: "https://google.com",
	},
	models.Link{
		Path: "/google-api",
		Url: "https://google.com/search?search=word",
	},
}

//Load : Implement
func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Link{}).Error
	if err != nil {
		log.Fatalf("Cannot drop table: &v", err)
	}
	err = db.Debug().AutoMigrate(&models.Link{}).Error
	if err != nil {
		log.Fatalf("Cannot migrate to database")
	}

	for i, _ := range links {
		err = db.Debug().Model(&models.Link{}).Create(&links[i]).Error
		if err != nil {
			log.Fatalf("Cannot seed links table: %v", err)
		}
	}
}