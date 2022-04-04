package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	perTime = 5
)

type website struct{
	name string
	url string
	status int 
}

var channel = make (chan website)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
		printStatus()
	}
}


func main() {
	doEvery(perTime*time.Second, checkUrlLoop)
}

func checkUrlLoop(t time.Time) {
	fmt.Println("Checking urls...")
	urls := returnUrls()

	for _, website := range urls {
		go checkUrl(website)
	}
}

func checkUrl(website website) {
	response, err := http.Get(website.url)

	if err != nil {
		website.status = -1
	} else {
		website.status = response.StatusCode
	}

	channel <- website
}

func returnUrls() []website {

	//we can remove this later
	urls := make(map[string]string)

	urls["golang"] = "https://go.dev/"
	urls["php"] = "https://www.php.net/"
	urls["laravel"] = "https://laravel.com/"
	urls["nodejs"] = "https://nodejs.org/en/"
	urls["mysql"] = "https://www.mysql.com/"
	urls["python"] = "https://www.python.org/"
	urls["postgres"] = "https://www.postgresql.org/"
	urls["haskel"] = "https://www.haskel.com/en-us"


	websites := []website{}

	for name, url := range urls{
		websites = append(websites, website{name, url, 0})
	}

	return websites
}

func printStatus(){
	for i := 0; i < len(returnUrls()); i++ {
		website := <- channel
		if website.status / 100 == 2 {
			fmt.Println("✅ ", website.name)
		}else{
			fmt.Println("❌", website.name)
		}
	}
}
