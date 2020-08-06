package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/", cardTest)
	r.GET("/", cardTest)
	r.GET("/haksik", hacksik)
	r.POST("/haksik", hacksik)
	r.GET("/library", library)
	r.POST("/library", library)
	r.GET("/information", information)
	r.POST("/information", information)
	r.GET("/weather", weather)
	r.POST("/weather", weather)
	r.Static("/assets", "./assets")
	r.Run("0.0.0.0:9091") // listen and serve on 0.0.0.0:8080
}
