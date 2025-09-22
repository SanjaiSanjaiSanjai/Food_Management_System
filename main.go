package main

import (
	baseurl "Food_Delivery_Management/BaseUrl"
	db "Food_Delivery_Management/DB"
	login "Food_Delivery_Management/Login"
	register "Food_Delivery_Management/Register"
	schema "Food_Delivery_Management/Schema"
	"Food_Delivery_Management/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	//call db config function to db connection
	db.DbConfig()

	//db.DB is not Empty success return function
	utils.IsNotNilSuccess(db.DB, "main", "Database is ready to use")

	// migrate schemas function
	schema.SchemaMigration()

	// get all routers in gin
	router := gin.Default()

	// return base URL ApiGroup function
	baseRoutes := baseurl.ApiGroup(router)

	// pass base URL RegisterRouter
	register.RegisterRouter(baseRoutes)

	// pass base URL LoginRoutes
	login.LoginRouter(baseRoutes)

	router.Run(":8081")
}
