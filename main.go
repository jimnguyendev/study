package main

import (
	"github.com/gin-gonic/gin"
	"github.com/study/common/uploadprovider"
	"github.com/study/component"
	"github.com/study/middleware"
	"github.com/study/modules/restaurant/restauranttransport/ginrestaurant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("DB")
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, s3Provider); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider) error {
	appCtx := component.NewAppContext(db, upProvider)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	// CRUD
	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	return r.Run()
}
