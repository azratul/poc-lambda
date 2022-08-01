package main

import (
	"lambda-parent/config"
	"log"
	"os"
)

func StartLocal() {
	if os.Getenv("API_URL1") == "" || os.Getenv("API_URL2") == "" {
		log.Println("All env vars must be set!!!")
		return
	}

	config.Conf.JsonPlaceHolder.UrlPosts = os.Getenv("API_URL1")
	config.Conf.JsonPlaceHolder.UrlUsers = os.Getenv("API_URL2")

	r := CreateServer()

	portEnv := os.Getenv("PORT")
	var port string
	if portEnv != "" {
		port = ":" + portEnv
	} else {
		port = ":8084"
	}

	r.Run(port)
}
