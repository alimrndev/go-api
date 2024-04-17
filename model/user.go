package model

import (
	"time"

	"github.com/alimrndev/go-api/config"
	"golang.org/x/crypto/bcrypt"
)

// User adalah struktur untuk representasi data pengguna
type User struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Username    string    `json:"username" gorm:"unique"`
	Password    string    `json:"-"` // "-" untuk menyembunyikan field ini dari serialisasi JSON
	Email       string    `json:"email" gorm:"unique"`
	GoogleID    string    `json:"google_id"`
	Role        string    `json:"role" gorm:"default:'customer'"`
	Fullname    string    `json:"fullname"`
	Picture     string    `json:"picture"`
	Gender      string    `json:"gender"`
	Birthdate   time.Time `json:"birthdate"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	CreatedBy   uint      `json:"created_by" gorm:"default:1"`
	UpdatedBy   uint      `json:"updated_by" gorm:"default:1"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

// SetPassword mengenkripsi password menggunakan bcrypt sebelum disimpan ke database
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword memeriksa apakah password yang diberikan cocok dengan password yang disimpan
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) Update() error {
	return config.DB.Save(&u).Error
}

func (u *User) Delete() error {
	return config.DB.Delete(&u).Error
}
