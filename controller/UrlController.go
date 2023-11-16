package controller

import (
	"fmt"
	"net/http"
	"time"
	"url_shortner/constant"
	"url_shortner/database"
	"url_shortner/helper"
	"url_shortner/types"

	"github.com/gin-gonic/gin"
)

// short the url
func ShortTheURL(c *gin.Context) {
	var shortUrlBody types.ShortUrlBody
	err := c.BindJSON(&shortUrlBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": constant.BindError})
		return
	}
	code := helper.GenRandomString(6)

	record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)

	if record.UrlCode != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "mesage": "this code is already in use"})
		return
	}

	var url types.URLDB

	url.CreatedAt = time.Now().Unix()
	url.ExpiredAt = time.Now().Unix()
	url.UrlCode = code
	url.LongUrl = shortUrlBody.LongUrl
	url.ShortUrl = constant.BaseUrl + code

	resp, err := database.Mgr.Insert(url, constant.UrlCollection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": false, "data": resp, "short_url": url.ShortUrl})
}

// redirect
func RedirectURL(c *gin.Context) {

	code := c.Param("code")

	record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)
	if record.UrlCode != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "meaasge": "there was no url found"})
	}
	fmt.Println(record.LongUrl)
}
