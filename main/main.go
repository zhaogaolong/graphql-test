package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"golang.zhaogaolong.com/graphql-test/api"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write(api.Page)
}

// nolint: gas
func main() {
	s, err := getSchema("./schema.gql")
	if err != nil {
		panic(err)
	}

	schema := graphql.MustParseSchema(s, &api.Resolver{})

	http.Handle("/", http.HandlerFunc(index))

	http.Handle("/query", &relay.Handler{Schema: schema})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
