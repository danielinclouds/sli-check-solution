package main

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealth(t *testing.T) {
	testCases := []struct {
		desc       string
		path       string
		wantStatus int
		wantBody   string
	}{
		{
			desc:       "One with bad path",
			path:       "http://127.0.0.1:8080/",
			wantStatus: 404,
			wantBody:   "Not Found",
		},
		{
			desc:       "One with correct path",
			path:       "http://127.0.0.1:8080/health",
			wantStatus: 200,
			wantBody:   "ok",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req := httptest.NewRequest("GET", tC.path, nil)
			w := httptest.NewRecorder()
			health(w, req)

			resp := w.Result()
			body, _ := ioutil.ReadAll(resp.Body)

			if resp.StatusCode != tC.wantStatus {
				t.Fatalf("Wrong status code, want: %v got: %v", tC.wantStatus, resp.StatusCode)
			}

			if strings.TrimSpace(string(body)) != tC.wantBody {
				t.Fatalf("Wrong body, want: %s got: %s", tC.wantBody, body)
			}
		})
	}
}
