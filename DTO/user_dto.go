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
	Postalcode int    `json:"postalcode" binding:"required"`
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
