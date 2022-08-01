package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"lambda-child02/config"
	"lambda-child02/jsonplaceholder"

	"log"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	log.Println("***GetUsers***")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d", config.Conf.JsonPlaceHolder.Url, id)
	var users jsonplaceholder.User

	if err := request(url, &users); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, users)
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
