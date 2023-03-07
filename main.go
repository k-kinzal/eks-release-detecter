package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const URL = "https://docs.aws.amazon.com/eks/latest/userguide/kubernetes-versions.html"

func main() {
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("ul").Each(func(i int, s *goquery.Selection) {
		if i != 0 {
			return
		}
		s.Find("p").Each(func(_ int, s *goquery.Selection) {
			fmt.Printf("%s\n", s.Text())
		})
	})
}
