package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	perTime = 5
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}


func main() {
	
	doEvery(perTime*time.Second, checkUrlLoop)
}

func checkUrlLoop(t time.Time) {
	fmt.Println("Checking urls...")
	urls := returnUrls()

	for name, url := range urls {
		go checkUrl(url, name)
	}
}

func checkUrl(url string, name string) {
	response, err := http.Get(url)

	if err != nil || response.StatusCode%2 != 0 {
		fmt.Println(name + " is Down")
	} else {
		fmt.Println(name + " is up")
	}
}

func returnUrls() map[string]string {
	urls := make(map[string]string)

	urls["golang"] = "https://go.dev/"
	urls["php"] = "https://www.php.net/"
	urls["laravel"] = "https://laravel.com/"
	urls["nodejs"] = "https://nodejs.org/en/"
	urls["mysql"] = "https://www.mysql.com/"
	urls["python"] = "https://www.python.org/"
	urls["postgres"] = "https://www.postgresql.org/"
	urls["haskel"] = "https://www.haskel.com/en-us"

	return urls
}
