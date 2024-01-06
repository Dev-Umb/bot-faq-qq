package store

import "github.com/jinzhu/gorm"

type (
	Factory interface {
	}
	DataStore struct {
		DB *gorm.DB
	}
)
