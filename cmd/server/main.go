package main

import (
	"github.com/IcaroTARique/RamenGoRV/config"
	"github.com/IcaroTARique/RamenGoRV/internal/entity"
	"github.com/IcaroTARique/RamenGoRV/internal/infra/database"
	"github.com/IcaroTARique/RamenGoRV/internal/infra/webserver/handlers"
	"github.com/IcaroTARique/RamenGoRV/internal/middleware"
	"github.com/go-chi/chi"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
)

// @title 			LamenGo API
// @version 		1.0.0
// @description 	API Test for RED|VENTURES
// @termsOfService 	http://swagger.io/terms/

// @contact.name 	Icaro Targino
// @contact.url 	http://www.icarotarginodev.com
// @contact.email	icaro_rique123@hotmail.com

// @license.name 	Icaro Targino Solucoes em TI
// @license.url 	http://www.icarotarginodev.com

// @host 			localhost:8080
// @BasePath 		/
// @securityDefinitions.apikey ApiKeyAuth
// @in 				header
// @name 			Authorization
func main() {
	configurations, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	dsn := configurations.DBUser + ":" + configurations.DBPassword + "@tcp(" + configurations.DBHost + ":" + configurations.DBPort + ")/" + configurations.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	//db, err := gorm.Open(mysql.Open("ramen.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&entity.Broth{}); err != nil {
		panic(err)
	}
	productDB := database.NewProduct(db)
	//productAPI := api.NewClient()

	productHandler := handlers.NewProductHandler(productDB)
	//productHandler := handlers.NewProductHandler(productAPI)

	r := chi.NewRouter()

	r.Use(middleware.ApiKeyVerifier)

	r.Get("/broths", productHandler.ListBroths)
	r.Get("/proteins", productHandler.ListProteins)
	r.Post("/order", productHandler.CreateOrder)

	if err := http.ListenAndServe(":"+configurations.WebServerPort, r); err != nil {
		panic(err)
	}
}
