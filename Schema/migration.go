package schema

import (
	db "Food_Delivery_Management/DB"
	enum "Food_Delivery_Management/ENUM"
	"Food_Delivery_Management/utils"
)

var exists bool

func SchemaMigration() {
	// if enum is already created return true
	db.DB.Raw(enum.EXISTS).Scan(&exists)

	// if exists is true db.DB.Exec is not working
	if !exists {
		db.DB.Exec(enum.ROLE)
	}

	//Migrate all Schemas (Table)
	err := db.DB.AutoMigrate(&User{}, &User_Addresses{}, &Role{}, &Restaurants{})

	// err is not nil return error
	utils.IsNotNilError(err, "SchemaMigration", "AutoMigration is issue")

	// err is  nil return Success
	utils.IsNillSuccess(err, "SchemaMigration", "AutoMigration is Success")
}
