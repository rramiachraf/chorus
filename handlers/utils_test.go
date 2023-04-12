package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	r := httptest.NewRecorder()
	data := map[string]string{"go": "test"}
	WriteJSON(r, 200, data)

	encoded, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	encoded = bytes.Join([][]byte{encoded, []byte("\n")}, []byte{})
	res := r.Result()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(body, encoded) != 0 {
		t.Fatalf("%q does not match %q", body, encoded)
	}
}

func TestHandleError(t *testing.T) {
	r := httptest.NewRecorder()
	status := http.StatusForbidden
	display := "permission denied"
	HandleError(r, status, nil, display)

	res := r.Result()

	if res.StatusCode != status {
		t.Fatalf("expected %d got %d", status, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	expectedBody := fmt.Sprintf("{\"error\":\"%s\"}\n", display)

	if string(body) != expectedBody {
		t.Fatalf("expected %q got %q", expectedBody, body)
	}
}

func TestNotFoundHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/not-found", nil)
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(NotFoundHandler)

	handler.ServeHTTP(rec, req)

	res := rec.Result()

	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("expected status code %d got %d", http.StatusNotFound, res.StatusCode)
	}
}
