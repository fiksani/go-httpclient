package main

import (
	"fmt"
	"github.com/fiksani/go-httpclient/gohttp"
	"io"
)

var (
	githubHttpClient = gohttp.New()
)

func getGithubHttpClient() gohttp.HttpClient {
	client := gohttp.New()

	//commonHeaders := make(http.Header)
	//commonHeaders.Set("Authorization", "Bearer abc123")
	//
	//client.SetHeaders(commonHeaders)

	return client
}

func main() {
	getUrls()
	getUrls()
	getUrls()
	getUrls()
}

func getUrls() {
	//headers := make(http.Header)
	//headers.Set("Authorization", "Bearer abc123")

	githubHttpClient := getGithubHttpClient()

	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
