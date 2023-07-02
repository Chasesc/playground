package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"urlshort"

	bolt "go.etcd.io/bbolt"
)

func main() {
	// Example of using the urlshort module.
	mux := defaultMux()

	db, err := bolt.Open("bolt.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(urlshort.BOLT_BUCKET))
		if err != nil {
			log.Fatal(err)
		}
		if err := bucket.Put([]byte("/go"), []byte("https://go.dev/")); err != nil {
			log.Fatal(err)
		}

		if err := bucket.Put([]byte("/vitess"), []byte("https://vitess.io/")); err != nil {
			log.Fatal(err)
		}
		return nil
	})

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/nearapp":        "https://nearapp.xyz",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml, err := os.ReadFile("urlmap.yaml")
	if err != nil {
		log.Fatal(err)
	}

	jsonBlob, err := os.ReadFile("urlmap.json")
	if err != nil {
		log.Fatal(err)
	}

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	jsonHandler, err := urlshort.JSONHandler([]byte(jsonBlob), yamlHandler)
	if err != nil {
		panic(err)
	}

	boltHandler := urlshort.BoltHandler(db, jsonHandler)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", boltHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
