package schema

import (
	"time"
)

// Schemas (Table)

type User struct {
	ID         uint      `gorm:"primaryKey;column:id"`
	Username   string    `gorm:"column:username;type:varchar(50);not null;uniqueIndex"`
	Email      string    `gorm:"column:email;type:varchar(255);not null;uniqueIndex"`
	Password   string    `gorm:"column:password;type:varchar(255);not null"`
	IsVerified bool      `gorm:"column:is_verified;default:false"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
	Status     bool      `gorm:"column:status;default:true"`

	// Relationships
	User_Addresses *User_Addresses `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Role           *Role           `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type User_Addresses struct {
	ID         uint      `gorm:"primaryKey;column:id"`
	UserID     uint      `gorm:"column:user_id;not null;index"`
	Address    string    `gorm:"column:address;type:text;not null"`
	State      string    `gorm:"column:state;type:varchar(100);not null"`
	Country    string    `gorm:"column:country;type:varchar(100);not null"`
	Postalcode string    `gorm:"column:postal_code;type:varchar(20);"`
	Landmark   string    `gorm:"column:landmark;type:varchar(255)"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
	Status     bool      `gorm:"column:status;default:true"`

	// Relationships
	User User `gorm:"foreignKey:UserID"`
}

type Role struct {
	ID        uint      `gorm:"primaryKey;column:id"`
	UserID    uint      `gorm:"column:user_id;not null;index"`
	Role      string    `gorm:"column:role;type:role_enum;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
	Status    bool      `gorm:"column:status;default:true"`

	// Relationships
	User *User `gorm:"foreignKey:UserID"`
}

type Restaurants struct {
	Id             uint      `gorm:"primaryKey;column:id"`
	Name           string    `gorm:"column:name;type:varchar(255);not null"`
	Description    string    `gorm:"column:description;type:text"`
	Rating         float64   `gorm:"column:rating;type:decimal(3,1);default:0.0"`
	Cuisine_type   string    `gorm:"column:cuisine_type;type:varchar(100)"`
	Phone          string    `gorm:"column:phone;type:varchar(20);not null"`
	Email          string    `gorm:"column:email;type:varchar(255);not null;unique"`
	License_number uint      `gorm:"column:license_number;not null;unique"`
	Owner_id       uint      `gorm:"column:owner_id;not null;index"`
	Status         bool      `gorm:"column:status;default:true"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime"`

	// Relationships
	Owner             *User              `gorm:"foreignKey:Owner_id;references:ID"`
	RestaurantAddress *RestaurantAddress `gorm:"foreignKey:RestaurantID"`
	Categories        []MenuCategory     `gorm:"many2many:restaurant_categories;foreignKey:Id;joinForeignKey:restaurant_id;References:ID;joinReferences:menu_category_id"`
}

type RestaurantAddress struct {
	ID           uint      `gorm:"primaryKey;column:id"`
	RestaurantID uint      `gorm:"column:restaurant_id;not null;index"`
	Address      string    `gorm:"column:address;type:text;not null"`
	City         string    `gorm:"column:city;type:varchar(100);not null"`
	State        string    `gorm:"column:state;type:varchar(100);not null"`
	PostalCode   string    `gorm:"column:postal_code;type:varchar(20)"`
	Latitude     float64   `gorm:"column:latitude;type:decimal(10,8)"`
	Longitude    float64   `gorm:"column:longitude;type:decimal(11,8)"`
	Status       bool      `gorm:"column:status;default:true"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`

	// Relationship
	Restaurant *Restaurants `gorm:"foreignKey:RestaurantID;references:Id"`
}

type MenuCategory struct {
	ID           uint      `gorm:"primaryKey;column:id"`
	Name         string    `gorm:"column:name;type:varchar(100);not null"`
	Description  string    `gorm:"column:description;type:text"`
	IsActive     bool      `gorm:"column:is_active;default:true;index:idx_menu_category_restaurant_active"`
	RestaurantID uint      `gorm:"column:restaurant_id;not null;index:idx_menu_category_restaurant,priority:1;index:idx_menu_category_restaurant_active,priority:1"`
	Status       bool      `gorm:"column:status;default:true;index:,type:btree"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`

	// Relationships
	Restaurant  *Restaurants  `gorm:"foreignKey:restaurant_id;references:id"`
	Menus       []Menu        `gorm:"foreignKey:CategoryID;references:ID"`
	Restaurants []Restaurants `gorm:"many2many:restaurant_categories;foreignKey:id;joinForeignKey:menu_category_id;References:id;joinReferences:restaurant_id"`
}

type Menu struct {
	ID           uint      `gorm:"primaryKey;column:id"`
	Name         string    `gorm:"column:name;type:varchar(255);not null"`
	Description  string    `gorm:"column:description;type:text"`
	Price        float64   `gorm:"column:price;type:decimal(10,2);not null"`
	IsVegetarian bool      `gorm:"column:is_vegetarian;default:false"`
	IsAvailable  bool      `gorm:"column:is_available;default:true"`
	RestaurantID uint      `gorm:"column:restaurant_id;not null;index"`
	CategoryID   uint      `gorm:"column:category_id;not null;index"`
	Status       bool      `gorm:"column:status;default:true"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`

	// Relationships
	Restaurant *Restaurants  `gorm:"foreignKey:RestaurantID;references:Id"`
	Category   *MenuCategory `gorm:"foreignKey:CategoryID;references:ID"`
}

type RestaurantCategory struct {
	ID           uint      `gorm:"primaryKey;column:id"`
	RestaurantID uint      `gorm:"column:restaurant_id;not null;index:idx_restaurant_category_restaurant,priority:1"`
	CategoryID   uint      `gorm:"column:menu_category_id;not null;index:idx_restaurant_category_category,priority:2"`
	Status       bool      `gorm:"column:status;default:true;index:,type:btree"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`

	// Relationships
	Restaurant Restaurants  `gorm:"foreignKey:restaurant_id;references:id"`
	Category   MenuCategory `gorm:"foreignKey:menu_category_id;references:id"`
}
