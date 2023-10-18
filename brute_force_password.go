package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	url := "http://ip/question1/"
	maxRequests := 1000
	waitTime := 0

	for i := 0; i < maxRequests; i++ {
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Status code:", resp.StatusCode)
		fmt.Println("Response Body:", string(body))

		resp.Body.Close()

		elapsed := time.Since(start)
		if elapsed > time.Second*10 {
			waitTime = int(elapsed.Seconds())
			break
		}
	}

	fmt.Printf("Wait time after hitting the limit is approximately %d seconds\n", (waitTime/10)*10)
}
