package models

import (
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

//Link :
type Link struct {
	Path string `gorm:"size:255;not null;`
	Url  string `gorm:"size:255;not null;`
}
//Prepare : Implement
func (l *Link) Prepare() {
	l.Path = html.EscapeString(strings.TrimSpace(l.Path))
	l.Url = html.EscapeString(strings.TrimSpace(l.Url))
}

//Save :
func (l *Link) Save(db *gorm.DB) (*Link, error) {
	var err error
	err = db.Debug().Model(&Link{}).Create(&l).Error
	if err != nil {
		return &Link{}, err
	}

	return l, nil
}

//FindAll : Implement
func (l *Link) FindAll(db *gorm.DB) (*[]Link, error) {
	var err error
	links := []Link{}

	err = db.Debug().Model(&Link{}).Limit(100).Find(&links).Error
	if err != nil {
		return &[]Link{}, err
	}
	return &links, nil
}