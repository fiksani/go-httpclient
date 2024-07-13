package main

import (
	"fmt"
	"github.com/fiksani/go-httpclient/gohttp"
	"io"
)

var (
	githubHttpClient = getGithubHttpClient()
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
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getUrls() {
	//headers := make(http.Header)
	//headers.Set("Authorization", "Bearer abc123")

	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func createUser(user User) {
	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
