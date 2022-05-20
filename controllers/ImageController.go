package controllers

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	imageNames []string
	imageQueue []string
)

func RandomImage(c *gin.Context) {
	imagePath := viper.GetString("Image.Path")

	if len(imageNames) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "no images found"})
		return
	}

	if len(imageQueue) == 0 {
		imageQueue = make([]string, len(imageNames))
		copy(imageQueue, imageNames)
	}

	img, idx := randImg(imageQueue)
	imageQueue = remove(imageQueue, idx)

	c.File(imagePath + "/" + img)
}

func ImageScanner() {
	fmt.Println("[ImageScanner] Scanning for images...")

	imagePath := viper.GetString("Image.Path")
	files, err := ioutil.ReadDir(imagePath)

	if err != nil {
		panic(err)
	}

	if len(files) == len(imageNames) {
		fmt.Println("[ImageScanner] Found no images")
		return
	}

	for _, file := range files {
		// if the image is not in the map, add it
		if !contains(imageNames, file.Name()) {
			imageNames = append(imageNames, file.Name())
		}
	}

	fmt.Println("[ImageScanner] Found " + strconv.Itoa(len(imageNames)) + " images")

	time.AfterFunc(time.Duration(viper.GetInt("Image.Scan.Interval"))*time.Minute, ImageScanner)
}

func randImg(arr []string) (string, int) {
	idx := rand.Intn(len(arr))
	return arr[idx], idx
}

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
