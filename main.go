package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	elasticsearch "github.com/elastic/go-elasticsearch"
	"github.com/gorilla/mux"
)

// HealthCheckHandler HealthCheck Endpoint
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}

// SearchHandler Search for a term
func SearchHandler(w http.ResponseWriter, r *http.Request) {

	var searchTerms Search

	if r.Body == nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&searchTerms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	results := SearchQuery(searchTerms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(results)
	if err != nil {
		log.Println(err)
	}

	io.WriteString(w, string(jsonData))

}

// SearchQuery provides wrapper for Elastic calls
func SearchQuery(search Search) (results Results) {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Search(
		es.Search.WithIndex("persons"),
		es.Search.WithQuery(search.Search),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	log.Println(res)

	var resultMap map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&resultMap); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	results.Count = int(resultMap["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	log.Printf("got total of %d", results.Count)

	return
}

// IndexHandler Index documents
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	jsonSearchDocument, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("No data to index")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var searchDocumentToIndex SearchDocument

	//put payload into a searchDocument
	json.Unmarshal(jsonSearchDocument, &searchDocumentToIndex)

	err = documentIndexer(searchDocumentToIndex)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	io.WriteString(w, `{"indexed": true}`)
}

//documentIndexer
func documentIndexer(s SearchDocument) error {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return errors.New("Error creating ES client")
	}

	data, _ := json.Marshal(s)

	res, err := es.Index(
		"persons",                       // Index name
		strings.NewReader(string(data)), // Document body
		es.Index.WithDocumentID(s.ID),   // Document ID
		es.Index.WithRefresh("true"),    // Refresh
	)
	if err != nil {
		return errors.New("Error indexing document")
	}
	defer res.Body.Close()
	return nil
}

//createIndex
func createIndex() error {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return errors.New("Error creating ES client")
	}

	es.Indices.Create("persons")

	return nil

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler)
	r.HandleFunc("/search", SearchHandler)
	r.HandleFunc("/index", IndexHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
