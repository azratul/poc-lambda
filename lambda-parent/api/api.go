package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"lambda-parent/config"
	"lambda-parent/jsonplaceholder"

	"log"

	ri "github.com/azratul/lambda-aws/rest_invoke"
	"github.com/gin-gonic/gin"
)

func GetPostsAndUsersTest(c *gin.Context) {
	log.Println("***GetPostsAndUsersTest***")
	var posts jsonplaceholder.Posts
	posts.Stage = os.Getenv("GIN_MODE")

	var lambda ri.Lambda
	lambda.Region = config.Conf.AWS.Region
	lambda.Function = config.Conf.AWS.LambdaPosts
	lambda.Object = &posts.Post
	lambda.Payload = ri.Payload{
		Path:       "/api/v1/posts",
		HTTPMethod: "GET",
	}

	err := lambda.RestInvoke()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, posts)
	return
}

func GetPostsAndUsers(c *gin.Context) {
	log.Println("***GetPostsAndUsers***")
	url := fmt.Sprintf("%s/posts", config.Conf.JsonPlaceHolder.UrlPosts)
	var posts jsonplaceholder.Posts
	posts.Stage = os.Getenv("GIN_MODE")

	if err := request(url, &posts.Post); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Fatal(err)
		return
	}

	var tmp []jsonplaceholder.User

	// TODO: Optimize this. Too many request. Save the data in a slice
	for i, post := range posts.Post {
		b, j := contains(&tmp, post.UserID)
		if b {
			posts.Post[i].User = tmp[j]
			continue
		}

		posts.Post[i].User = jsonplaceholder.User{}
		if err := request(fmt.Sprintf("%s/users/%d", config.Conf.JsonPlaceHolder.UrlUsers, post.UserID), &posts.Post[i].User); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			log.Fatal(err)
			return
		}

		tmp = append(tmp, posts.Post[i].User)
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

func contains(s *[]jsonplaceholder.User, e int) (bool, int) {
	for i, a := range *s {
		if a.ID == e {
			return true, i
		}
	}
	return false, -1
}
