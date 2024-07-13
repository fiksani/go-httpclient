package main

import (
	"fmt"
	"github.com/fiksani/go-httpclient/gohttp"
	"io"
)

func main() {
	client := gohttp.New()

	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
