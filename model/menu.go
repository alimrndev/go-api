package model

import (
	"time"

	"github.com/alimrndev/go-api/config"
)

// User adalah struktur untuk representasi data pengguna
type MenuItem struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	Name          string    `json:"name"`
	Description   string    `json:"description" gorm:"unique"`
	Picture       string    `json:"picture"`
	Price         float32   `json:"price"`
	Qty           string    `json:"qty"`
	DiscountType  string    `json:"discount_type" gorm:"default:'nominal'"`
	DiscountValue float32   `json:"discount_value" gorm:"default:0"`
	Category      string    `json:"category"`
	CreatedBy     uint      `json:"created_by" gorm:"default:1"`
	UpdatedBy     uint      `json:"updated_by" gorm:"default:1"`
	CreatedAt     time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (m *MenuItem) Update() error {
	return config.DB.Save(&m).Error
}

func (m *MenuItem) Delete() error {
	return config.DB.Delete(&m).Error
}
