package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type jsonData struct {
	input  interface{}
	output string
}

func TestWriteJSON(t *testing.T) {
	tests := []jsonData{
		{map[string]string{"go": "test"}, `{"go":"test"}` + "\n"},
		{map[string]int{"age": 100}, `{"age":100}` + "\n"},
	}

	for _, test := range tests {
		r := httptest.NewRecorder()
		WriteJSON(r, 200, test.input)

		res := r.Result()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		if string(body) != test.output {
			t.Fatalf("%q does not match %q", body, test.output)
		}

	}
}

type testErrors struct {
	statusCode int
	display    string
}

func TestHandleError(t *testing.T) {
	tests := []testErrors{
		{http.StatusForbidden, "permission denied"},
		{http.StatusInternalServerError, "something went wrong"},
	}

	for _, test := range tests {
		r := httptest.NewRecorder()
		HandleError(r, test.statusCode, nil, test.display)

		res := r.Result()
		if res.StatusCode != test.statusCode {
			t.Fatalf("expected %d got %d", test.statusCode, res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		expectedBody := fmt.Sprintf(`{"error":"%s"}`+"\n", test.display)

		if string(body) != expectedBody {
			t.Fatalf("expected %q got %q", expectedBody, body)
		}
	}
}

type notFoundRequest struct {
	method string
	path   string
}

func TestNotFoundHandler(t *testing.T) {
	tests := []notFoundRequest{
		{http.MethodGet, "/blablabla"},
		{http.MethodPost, "/not-found"},
		{http.MethodDelete, "/does-not-exist/does-not-exist-too"},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.path, nil)
		rec := httptest.NewRecorder()
		handler := http.HandlerFunc(NotFoundHandler)

		handler.ServeHTTP(rec, req)

		res := rec.Result()

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("expected status code %d got %d", http.StatusNotFound, res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		expectedErr := `{"error":"endpoint not found"}` + "\n"
		if string(body) != expectedErr {
			t.Fatalf("expected response %q got %q", expectedErr, body)
		}
	}
}
