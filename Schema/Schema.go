package schema

import (
	"time"
)

// Schemas (Table)

type User struct {
	ID             uint   `gorm:"primaryKey"`
	Username       string `gorm:"uniqueIndex"`
	Email          string `gorm:"uniqueIndex"`
	Password       string
	IsVerified     bool `gorm:"column:is_verified;default:false"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Status         bool            `gorm:"default:true"`
	User_Addresses *User_Addresses `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Role           *Role           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type User_Addresses struct {
	Id         uint `gorm:"primaryKey"`
	UserID     uint `gorm:"not null"`
	Address    string
	State      string
	Country    string
	Postalcode int
	Landmark   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Status     bool
	User       User `gorm:"foreignKey:UserID"`
}

type Role struct {
	Id          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	Role        string `gorm:"type:role_enum"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Status      bool
	User        User          `gorm:"foreignKey:UserID"`
	Restaurants []Restaurants `gorm:"foreignKey:Owner_id"`
}

type Restaurants struct {
	Id             uint `gorm:"primaryKey"`
	Name           string
	Description    string
	Rating         float64
	Cuisine_type   string
	Phone          string
	Email          string
	License_number uint
	Owner_id       uint
	Status         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time

	Role *Role `gorm:"foreignKey:Owner_id;references:Id"`
}
