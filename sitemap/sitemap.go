package sitemap // import "gophercise/sitemap"

import (
	"gophercise/link"
	"log"
	"net/http"
)

func mapSite(url string) {
	//open reader from the url
	r, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}

	// get links from that page
	log.Println(link.GetLinks(r.Body))
}
