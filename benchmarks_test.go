package main

import (
	"golb/graph"
	"golb/graph/generated"
	"golb/models"
	"golb/services"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
)

func BenchmarkSimpleQueryNoArgs(b *testing.B) {
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{},
	}))

	q := `{"query":"{ setting { id } }"}`

	var body strings.Reader
	r := httptest.NewRequest("POST", "/query", &body)
	r.Header.Set("Content-Type", "application/json")

	b.ReportAllocs()
	b.ResetTimer()

	rec := httptest.NewRecorder()
	for i := 0; i < b.N; i++ {
		body.Reset(q)
		rec.Body.Reset()
		server.ServeHTTP(rec, r)
		if rec.Body.String() != `{"data":{"setting":{"id":"1"}}}` {
			b.Fatalf("Unexpected response: %s", rec.Body.String())
		}
	}
}

func BenchmarkDatabase(b *testing.B) {
	tx := services.DB
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var setting models.Setting
		tx.First(&setting)
	}
}

func BenchmarkMultiQuery(b *testing.B) {
	server := httptest.NewServer(ginSetup())
	defer server.Close()

	q := `{"query": "{allSettings{settings{id}}allRoles{roles{id}}}"}`
	r := `{"data":{"allSettings":{"settings":[{"id":"1"},{"id":"2"},{"id":"3"},{"id":"4"},{"id":"5"},{"id":"6"},{"id":"7"}]},"allRoles":{"roles":[{"id":"1"},{"id":"2"},{"id":"3"},{"id":"4"},{"id":"5"}]}}}`

	var body strings.Reader

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		body.Reset(q)
		resp, err := http.Post(server.URL+"/query", "application/json", &body)
		if err != nil {
			b.Error("Encountered an error when processing request: ", err)
		}
		result, _ := ioutil.ReadAll(resp.Body)
		if string(result) != r {
			b.Fatalf("Unexpected response: %s", string(result))
		}
	}

}

func BenchmarkNestedQuery(b *testing.B) {
	server := httptest.NewServer(ginSetup())
	defer server.Close()

	q := `{"query": "{allUsers{users{id postConnection{posts{id authorConnection{authors{id name}}}}}}}"}` // 3 levels of nesting

	var body strings.Reader

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		body.Reset(q)
		if _, err := http.Post(server.URL+"/query", "application/json", &body); err != nil {
			b.Error("Encountered an error when processing request: ", err)
		}
	}
}

func BenchmarkNestedOverLimitedQuery(b *testing.B) {
	server := httptest.NewServer(ginSetup())
	defer server.Close()

	q := `{"query": "{allUsers{users{id postConnection{posts{id authorConnection{authors{id postConnection{posts{id authorConnection{authors{id}}}}}}}}}}}"}` // 3 levels of nesting
	r := `{"errors":[{"message":"operation has complexity 683, which exceeds the limit of 512","extensions":{"code":"COMPLEXITY_LIMIT_EXCEEDED"}}],"data":null}`
	var body strings.Reader

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		body.Reset(q)
		resp, err := http.Post(server.URL+"/query", "application/json", &body)
		if err != nil {
			b.Error("Encountered an error when processing request: ", err)
		}
		result, _ := ioutil.ReadAll(resp.Body)
		if string(result) != r {
			b.Fatalf("Unexpected response: %s", string(result))
		}
	}

}
