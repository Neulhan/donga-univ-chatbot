package main

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func cardTest(c *gin.Context) {
	c.JSON(200, textTemplate([]string{"111"}))
}

func hacksik(c *gin.Context) {
	campus := c.Query("campus")
	text := crawlingHaksik(campus)
	c.JSON(200, textTemplate([]string{text}))
}

func library(c *gin.Context) {
	text := crwalingLibrary()
	c.JSON(200, textTemplate([]string{text}))
}

func information(c *gin.Context) {
	text := crawlingInformation()
	c.JSON(200, textTemplate([]string{text}))
}

func weather(c *gin.Context) {
	var nx string
	var ny string
	var webLink string
	campus := c.Query("campus")
	switch campus {
	case "승학":
		// 하단제1동
		nx = "96"
		ny = "74"
		webLink = "https://search.naver.com/search.naver?sm=tab_hty.top&where=nexearch&query=%ED%95%98%EB%8B%A8%EC%A0%9C1%EB%8F%99+%EB%82%A0%EC%94%A8&oquery=%ED%95%98%EB%8B%A8%EC%A0%9C1%EB%8F%99&tqi=UYuhmwprvmsssk0wz8lsssssttd-411972"
		break
	case "부민":
		//부민동
		nx = "97"
		ny = "74"
		webLink = "https://search.naver.com/search.naver?sm=top_hty&fbm=1&ie=utf8&query=%EB%B6%80%EB%AF%BC%EB%8F%99+%EB%82%A0%EC%94%A8"
		break
	case "구덕":
		//동대신제1동
		nx = "97"
		ny = "74"
		webLink = "https://search.naver.com/search.naver?sm=tab_hty.top&where=nexearch&query=%EB%8F%99%EB%8C%80%EC%8B%A0%EC%A0%9C1%EB%8F%99+%EB%82%A0%EC%94%A8&oquery=%EB%B6%80%EB%AF%BC%EB%8F%99+%EB%82%A0%EC%94%A8&tqi=UYuhewprvmZsssQds2ZssssstdV-233946"
		break
	}
	imgURL, text := getWeatherData(nx, ny)
	button := buttonTemplate("webLink", "자세히보기", webLink)

	c.JSON(200, cardTemplate("지금 "+campus+"캠퍼스 날씨는", text, imgURL, []j{button}))
}
