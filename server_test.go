package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestQuery(t *testing.T) {

	var testCase = []struct {
		request, response string
	}{
		{`{"query":"{setting{id}}"}`, `{"data":{"setting":{"id":"1"}}}`},
		{`{"query":"{setting(id:2){id}}"}`, `{"data":{"setting":{"id":"2"}}}`},
	}

	for _, tC := range testCase {
		expect(tC.request, tC.response, t)
	}
}

func expect(req, res string, t *testing.T) {
	if !json.Valid([]byte(req)) || !json.Valid([]byte(res)) {
		t.Fatal("Invalid json string")
	}
	server := httptest.NewServer(ginSetup())
	defer server.Close()

	var body strings.Reader
	body.Reset(req)
	resp, err := http.Post(server.URL+"/query", "application/json", &body)
	if err != nil {
		t.Error("Encountered an error when processing request: ", err)
	}
	result, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("Encountered an error when transform response: ", err)
	}
	if string(result) != res {
		t.Error("Unexpected response: ", string(result))
	}
}
