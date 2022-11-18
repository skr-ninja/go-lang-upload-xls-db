package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skr/models"
	"github.com/skr/uploads"
)

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "You are at home",
	})
}

func routes() {

	models.ConnectDataBase()

	router := gin.Default()
	router.GET("/", home)
	router.POST("/upload", uploads.UploadFile)
	//router.POST("/get-files", getAllUploadedAssets)
	//router.GET("/get-upload/:assetId", getUploadedFile)
	//router.PUT("/update-file/:name", updateFile)
	//router.DELETE("/delete-file/:assetId", deleteAsset)
	// Default port is 8080 you can change it to any port you like
	//r.Run(":8080")
	router.Run(":8089")

}

func main() {
	routes()
}
