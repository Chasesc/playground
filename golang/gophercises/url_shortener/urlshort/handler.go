package urlshort

import (
	"encoding/json"
	"log"
	"net/http"

	bolt "go.etcd.io/bbolt"
	"gopkg.in/yaml.v3"
)

const BOLT_BUCKET = "URL_MAP"

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if url, ok := pathsToUrls[req.RequestURI]; ok {
			http.Redirect(w, req, url, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, req)
	}
}

type urlMapping struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url" json:"url"`
}

func createUrlMap(mapSequence []urlMapping) map[string]string {
	mappings := make(map[string]string)
	for _, urlMap := range mapSequence {
		mappings[urlMap.Path] = urlMap.Url
	}
	return mappings
}

func ymlToMap(yml []byte) (map[string]string, error) {
	var mapSequence []urlMapping
	if err := yaml.Unmarshal(yml, &mapSequence); err != nil {
		return nil, err
	}

	return createUrlMap(mapSequence), nil
}

func jsonToMap(jsonBlob []byte) (map[string]string, error) {
	var mapSequence []urlMapping
	if err := json.Unmarshal(jsonBlob, &mapSequence); err != nil {
		return nil, err
	}

	return createUrlMap(mapSequence), nil
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned are related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathToUrls, err := ymlToMap(yml)
	if err != nil {
		return nil, err
	}
	return MapHandler(pathToUrls, fallback), nil
}

func JSONHandler(jsonBlob []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathToUrls, err := jsonToMap(jsonBlob)
	if err != nil {
		return nil, err
	}
	return MapHandler(pathToUrls, fallback), nil
}

func BoltHandler(db *bolt.DB, fallback http.Handler) http.HandlerFunc {
	db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(BOLT_BUCKET)); err != nil {
			log.Fatal(err)
		}
		return nil
	})

	return func(w http.ResponseWriter, req *http.Request) {
		db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(BOLT_BUCKET))
			url := bucket.Get([]byte(req.RequestURI))

			if len(url) > 0 {
				http.Redirect(w, req, string(url), http.StatusFound)
			}

			return nil
		})

		fallback.ServeHTTP(w, req)
	}
}
