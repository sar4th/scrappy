package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/http"
)

func main() {
	scrapeHandler := func(w http.ResponseWriter, req *http.Request) {

		userURL := req.URL.Query().Get("url")
		if userURL == "" {
			http.Error(w, "Please provide a url to scrape", http.StatusBadRequest)
			return
		}

		c := colly.NewCollector()
		fmt.Println(userURL)
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			fmt.Println(link)
		})

		err := c.Visit(userURL)
		if err != nil {
			fmt.Println("Error visiting URL:", err)
		}
	}

	http.HandleFunc("/scrape", scrapeHandler)

	log.Println("Listening for requests at http://localhost:8000/scrape")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
