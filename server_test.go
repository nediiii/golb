package main

import (
	"golb/graph"
	"golb/graph/generated"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
)

func TestServer(t *testing.T) {
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{},
	}))

	q := `{"query":"{ setting { id } }"}`

	var body strings.Reader
	r := httptest.NewRequest("POST", "/query", &body)
	r.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	body.Reset(q)
	rec.Body.Reset()
	server.ServeHTTP(rec, r)
	if rec.Body.String() != `{"data":{"setting":{"id":"1"}}}` {
		t.Fatalf("Unexpected response: %s", rec.Body.String())
	}
}
