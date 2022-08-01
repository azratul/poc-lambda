package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"lambda-child01/config"
	"lambda-child01/jsonplaceholder"

	"log"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	log.Println("***GetPosts***")
	url := fmt.Sprintf("%s/posts", config.Conf.JsonPlaceHolder.Url)
	var posts []jsonplaceholder.Post

	if err := request(url, &posts); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, posts)
}

func request(url string, obj interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(&obj)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	req.Close = true

	return nil
}
