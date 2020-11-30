package wrapper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockJSONStructure struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

var w HTTPWrapper

func TestGET(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(w.GET(func(http.ResponseWriter, *http.Request) {}))
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf(methodNotAllowedErr, request.Method)
	}
}

func TestPOST(t *testing.T) {
	request, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(w.POST(func(http.ResponseWriter, *http.Request) {}))
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf(methodNotAllowedErr, request.Method)
	}
}

func TestPUT(t *testing.T) {
	request, err := http.NewRequest("PUT", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(w.PUT(func(http.ResponseWriter, *http.Request) {}))
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf(methodNotAllowedErr, request.Method)
	}
}

func TestDELETE(t *testing.T) {
	request, err := http.NewRequest("DELETE", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(w.DELETE(func(http.ResponseWriter, *http.Request) {}))
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf(methodNotAllowedErr, request.Method)
	}
}

func TestMethodGet(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(w.METHOD("GET", func(http.ResponseWriter, *http.Request) {}))
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf(methodNotAllowedErr, request.Method)
	}
}

func TestMethodPatch(t *testing.T) {
	request, err := http.NewRequest("PATCH", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(w.METHOD("PATCH", func(http.ResponseWriter, *http.Request) {}))
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf(methodNotAllowedErr, request.Method)
	}
}

func TestMethodsGetPost(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(w.METHODS([]string{"GET", "POST"}, func(http.ResponseWriter, *http.Request) {}))
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf(methodNotAllowedErr, request.Method)
	}
}

func TestJSONResponse(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(w.JSONResponse(w.GET(func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(MockJSONStructure{})
	})))
	handler.ServeHTTP(recorder, request)

	if recorder.Header().Get("Content-Type") != "application/json" {
		t.Error("Content-Type is not application/json")
	}
}
