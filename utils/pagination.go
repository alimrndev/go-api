package utils

import (
	"github.com/jinzhu/gorm"
)

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

func Paginate(db *gorm.DB, pagination Pagination) (*gorm.DB, error) {
	if pagination.Page == 0 {
		pagination.Page = 1
	}
	if pagination.PerPage == 0 {
		pagination.PerPage = 10 // Default perPage, bisa disesuaikan dengan kebutuhan
	}

	offset := (pagination.Page - 1) * pagination.PerPage
	db = db.Offset(offset).Limit(pagination.PerPage)
	return db, nil
}
