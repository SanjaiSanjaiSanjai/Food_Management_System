package DTO

type RegisterDTO struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required"`
}

type UserAddressDTO struct {
	Address    string `json:"address" binding:"required"`
	Country    string `json:"country" binding:"required"`
	State      string `json:"state" binding:"required"`
	Postalcode string `json:"postalcode" binding:"required"`
	Landmark   string `json:"landmark" binding:"required"`
	Status     bool   `json:"status" binding:"required"`
}

type RestaurantDTO struct {
	Name           string  `json:"name" binding:"required"`
	Description    string  `json:"description" binding:"required"`
	Rating         float64 `json:"rating" binding:"required"`
	Cuisine_type   string  `json:"cuisine_type" binding:"required"`
	Phone          string  `json:"phone" binding:"required"`
	Email          string  `json:"email" binding:"required,email"`
	License_number uint    `json:"license_number" binding:"required"`
}

type RestaurantAddressDTO struct {
	Address    string  `json:"address" binding:"required"`
	City       string  `json:"city" binding:"required"`
	State      string  `json:"state" binding:"required"`
	Postalcode string  `json:"postalcode" binding:"required"`
	Latitude   float64 `json:"latitude" binding:"required"`
	Longitude  float64 `json:"longitude" binding:"required"`
}

type MenuDTO struct {
	Name         string  `json:"name" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
	IsVegetarian bool    `json:"is_vegetarian" binding:"required"`
	CategoryName string  `json:"category_name" binding:"required"`
}

type MenuCategoryDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsActive    bool   `json:"is_active" binding:"required"`
}
