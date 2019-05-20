package testrequest

import (
	"encoding/json"
	"io"
	"log"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"app/api/component"
)

// SendForm is a helper to quickly make a form request.
func SendForm(t *testing.T, core component.Core, method string, target string,
	v url.Values) *httptest.ResponseRecorder {
	component.LoadRoutes(core)

	var body io.Reader
	if v != nil {
		body = strings.NewReader(v.Encode())
	}

	r := httptest.NewRequest(method, target, body)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	core.Router.ServeHTTP(w, r)

	return w
}

// SendJSON is a helper to quickly make a JSON request.
func SendJSON(t *testing.T, core component.Core, method string, target string,
	v url.Values) *httptest.ResponseRecorder {
	component.LoadRoutes(core)

	var body io.Reader
	if v != nil {
		body = strings.NewReader(ToJSON(v))
	}

	r := httptest.NewRequest(method, target, body)
	r.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	core.Router.ServeHTTP(w, r)

	return w
}

// ToJSON converts a url.Values to a JSON string.
func ToJSON(values url.Values) string {
	m := make(map[string]string)

	for k, v := range values {
		if len(v) > 0 {
			m[k] = v[0]
		} else {
			m[k] = ""
		}
	}

	js, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	return string(js)
}