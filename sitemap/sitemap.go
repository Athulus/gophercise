package sitemap // import "github.com/Athulus/gophercise/sitemap"

import (
	"log"
	"net/http"

	"github.com/Athulus/gophercise/link"
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
