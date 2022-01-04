package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// performGetRequest()
	// performPOSTRequest()
	performPOSTFormRequest()
}

func performGetRequest() {
	fmt.Println("This is a web verb code")
	const url string = "http://localhost:8000/get"

	res, err := http.Get(url)
	handleErr(err)
	defer res.Body.Close()

	fmt.Println("The status code: ", res.StatusCode)
	fmt.Println("The content length: ", res.ContentLength)

	var responseString strings.Builder
	content, err := ioutil.ReadAll(res.Body)
	handleErr(err)

	byteCount, err := responseString.Write(content)
	handleErr(err)

	fmt.Println(byteCount)
	fmt.Println("Content is ", responseString.String())
	fmt.Println("Content is ", responseString.Cap())
}

func performPOSTRequest() {
	const MyURL string = "http://localhost:8000/post"
	reqBody := strings.NewReader(`
		{
			"Google":"150k",
			"Netflix":"174k"
		}
	`)

	res, err := http.Post(MyURL, "application/json", reqBody)
	handleErr(err)
	defer res.Body.Close()

	var responseString strings.Builder
	content, err := ioutil.ReadAll(res.Body)
	handleErr(err)
	responseString.Write(content)
	fmt.Println(responseString.String())
}

func performPOSTFormRequest() {
	const MyURL string = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstName", "Aniruddha")
	data.Add("lastName", "Basak")
	data.Add("email", "aniruddha@google.com")

	res, err := http.PostForm(MyURL, data)
	handleErr(err)
	defer res.Body.Close()

	var responseString strings.Builder
	content, err := ioutil.ReadAll(res.Body)
	handleErr(err)
	responseString.Write(content)
	fmt.Println(responseString.String())

}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
