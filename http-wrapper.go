package wrapper

import (
	"fmt"
	"net/http"
	"strings"
)

// HTTPWrapper -
type HTTPWrapper struct{}

var methodNotAllowedErr = "Method %s is not allowed."

// GET -
func (r *HTTPWrapper) GET(f http.HandlerFunc) http.HandlerFunc {
	return handleMethod(f, http.MethodGet)
}

// POST -
func (r *HTTPWrapper) POST(f http.HandlerFunc) http.HandlerFunc {
	return handleMethod(f, http.MethodPost)
}

// PUT -
func (r *HTTPWrapper) PUT(f http.HandlerFunc) http.HandlerFunc {
	return handleMethod(f, http.MethodPut)
}

// DELETE -
func (r *HTTPWrapper) DELETE(f http.HandlerFunc) http.HandlerFunc {
	return handleMethod(f, http.MethodDelete)
}

// METHOD -
func (r *HTTPWrapper) METHOD(method string, f http.HandlerFunc) http.HandlerFunc {
	return handleMethod(f, strings.ToUpper(method))
}

// METHODS -
func (r *HTTPWrapper) METHODS(methods []string, f http.HandlerFunc) http.HandlerFunc {
	return handleMethod(f, methods...)
}

// handleMethod -
func handleMethod(f http.HandlerFunc, methods ...string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, method := range methods {
			if r.Method == method {
				f(w, r)
				return
			}
		}
		err := fmt.Sprintf(methodNotAllowedErr, r.Method)
		http.Error(w, err, http.StatusMethodNotAllowed)
	})
}

// JSONResponse -
func (r *HTTPWrapper) JSONResponse(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
