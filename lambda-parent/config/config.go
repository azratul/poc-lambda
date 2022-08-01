package config

var Conf Config

type Config struct {
	JsonPlaceHolder struct {
		UrlPosts string
		UrlUsers string
		ApiKey   string
	}
	AWS struct {
		LambdaPosts string
		LambdaUsers string
		Region      string
	}
}
