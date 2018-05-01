package urlshort // import "urlshort"

import (
	"log"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

/*
MapHandler will return an http HandlerFunc
that will redirect a url if it is in the passed in map
otherwise it will pass the request through to the fallback handler
*/
func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if newURL, ok := pathToUrls[r.URL.Path]; ok {
			log.Printf("redirecting to *%v*", newURL)
			http.Redirect(w, r, newURL, http.StatusFound)
		} else {
			log.Printf("could not find url *%v* in the redirect input ", r.URL.Path)
			fallback.ServeHTTP(w, r)
		}
	}
}

/*
YamlHandler will do the same thing as MapHandler but you pass it a yaml file
*/
func YamlHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	//parse yaml to map
	urls := make(map[string]string)
	var yamlList []pathURL
	err := yaml.Unmarshal(yml, &yamlList)
	if err != nil {
		return nil, err
	}

	// change yaml sturct into a map for MapHandler
	for _, pu := range yamlList {
		urls[pu.Path] = pu.URL
	}
	return MapHandler(urls, fallback), nil
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
