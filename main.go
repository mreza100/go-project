package main

import (
	"Waringin/controllers/aboutcontroller"
	"Waringin/controllers/articlecategoriescontroller"
	"Waringin/controllers/articlenewscontroller"
	"Waringin/controllers/brandcontroller"
	"Waringin/controllers/hotelcontroller"
	"Waringin/controllers/hotelfacilitiescontroller"
	"Waringin/controllers/hotelgallerycontroller"
	"Waringin/controllers/hotelinformationcontroller"
	"Waringin/controllers/hotellocationcontroller"
	"Waringin/controllers/hotelroomtypecontroller"
	"Waringin/controllers/managementinvestorcontroller"
	"Waringin/controllers/specialoffercontroller"
	authcontroller "Waringin/controllers/usercontroller"
	"Waringin/middlewares"
	"Waringin/models"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func main() {

	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://5cc150af29874728929f3fdb768cf2e2@arcus.microad.co.id/17",
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")

	r := gin.Default()
	models.ConnectDatabase()
	models.ConnectCDN()
	models.Seeder()

	//Auth
	r.Static("/upload", "./upload")
	r.POST("/api/login", authcontroller.Login)

	//Frontend Result
	r.GET("/resultbrand", brandcontroller.Result)
	r.GET("/resultabout", aboutcontroller.Result)
	r.GET("/getlocation", hotellocationcontroller.Index)
	r.GET("/resultlocation/:id", hotelcontroller.Result)

	//User
	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", authcontroller.CurrentUser)
	protected.POST("/register", authcontroller.Register)
	protected.GET("/getalluser", authcontroller.Index)
	protected.POST("/updateuser", authcontroller.Update)

	//Brands
	protected.GET("/brands", brandcontroller.Index)
	protected.GET("/brands/:id", brandcontroller.Show)
	protected.POST("/brands", brandcontroller.Create)
	protected.PUT("/brands/:id", brandcontroller.Update)
	protected.DELETE("/brands", brandcontroller.Delete)

	//About
	protected.GET("/abouts", aboutcontroller.Index)
	protected.GET("/abouts/:id", aboutcontroller.Show)
	protected.POST("/abouts", aboutcontroller.Create)
	protected.PUT("/abouts/:id", aboutcontroller.Update)
	protected.DELETE("/abouts", aboutcontroller.Delete)

	//Hotel Location
	protected.GET("/hotellocation", hotellocationcontroller.Index)
	protected.GET("/hotellocation/:id", hotellocationcontroller.Show)
	protected.POST("/hotellocation", hotellocationcontroller.Create)
	protected.PUT("/hotellocation/:id", hotellocationcontroller.Update)
	protected.DELETE("/hotellocation", hotellocationcontroller.Delete)

	//Special Offer
	protected.GET("/specialoffer", specialoffercontroller.Index)
	protected.GET("/specialoffer/:id", specialoffercontroller.Show)
	protected.POST("/specialoffer", specialoffercontroller.Create)
	protected.PUT("/specialoffer/:id", specialoffercontroller.Update)
	protected.DELETE("/specialoffer", specialoffercontroller.Delete)

	//Hotel
	protected.GET("/hotel", hotelcontroller.Index)
	protected.GET("/hotel/:id", hotelcontroller.Show)
	protected.POST("/hotel", hotelcontroller.Create)
	protected.PUT("/hotel/:id", hotelcontroller.Update)
	protected.DELETE("/hotel", hotelcontroller.Delete)

	//Hotel Room Type
	protected.GET("/hotelroomtype", hotelroomtypecontroller.Index)
	protected.GET("/hotelroomtype/:id", hotelroomtypecontroller.Show)
	protected.POST("/hotelroomtype", hotelroomtypecontroller.Create)
	protected.PUT("/hotelroomtype/:id", hotelroomtypecontroller.Update)
	protected.DELETE("/hotelroomtype", hotelroomtypecontroller.Delete)

	//Hotel Facilities
	protected.GET("/hotelfacilities", hotelfacilitiescontroller.Index)
	protected.GET("/hotelfacilities/:id", hotelfacilitiescontroller.Show)
	protected.POST("/hotelfacilities", hotelfacilitiescontroller.Create)
	protected.PUT("/hotelfacilities/:id", hotelfacilitiescontroller.Update)
	protected.DELETE("/hotelfacilities", hotelfacilitiescontroller.Delete)

	//Hotel Gallery
	protected.GET("/hotelgallery", hotelgallerycontroller.Index)
	protected.GET("/hotelgallery/:id", hotelgallerycontroller.Show)
	protected.POST("/hotelgallery", hotelgallerycontroller.Create)
	protected.PUT("/hotelgallery/:id", hotelgallerycontroller.Update)
	protected.DELETE("/hotelgallery", hotelgallerycontroller.Delete)

	//Hotel Information
	protected.GET("/hotelinformation", hotelinformationcontroller.Index)
	protected.GET("/hotelinformation/:id", hotelinformationcontroller.Show)
	protected.POST("/hotelinformation", hotelinformationcontroller.Create)
	protected.PUT("/hotelinformation/:id", hotelinformationcontroller.Update)
	protected.DELETE("/hotelinformation", hotelinformationcontroller.Delete)

	//Management Investor
	protected.GET("/managementinvestor", managementinvestorcontroller.Index)
	protected.GET("/managementinvestor/:id", managementinvestorcontroller.Show)
	protected.POST("/managementinvestor", managementinvestorcontroller.Create)
	protected.PUT("/managementinvestor/:id", managementinvestorcontroller.Update)
	protected.DELETE("/managementinvestor", managementinvestorcontroller.Delete)

	//Article News
	protected.GET("/articlenews", articlenewscontroller.Index)
	protected.GET("/articlenews/:id", articlenewscontroller.Show)
	protected.POST("/articlenews", articlenewscontroller.Create)
	protected.PUT("/articlenews/:id", articlenewscontroller.Update)
	protected.DELETE("/articlenews", articlenewscontroller.Delete)

	//Article Categories
	protected.GET("/articlecategories", articlecategoriescontroller.Index)
	protected.GET("/articlecategories/:id", articlecategoriescontroller.Show)
	protected.POST("/articlecategories", articlecategoriescontroller.Create)
	protected.PUT("/articlecategories/:id", articlecategoriescontroller.Update)
	protected.DELETE("/articlecategories", articlecategoriescontroller.Delete)

	r.Run(":6000")
}
