package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gospeak/auth-service/api/middleware"
	"github.com/gospeak/auth-service/api/server"
)

type Final struct {
	IsCall bool
	Next   server.HandlerFunc
}

func (f *Final) Handler(ctx server.Context, w http.ResponseWriter, r *http.Request) {
	f.IsCall = true
	f.Next(ctx, w, r)
}

func TestInitToken(t *testing.T) {

	type testcase struct {
		ReqGen           func(m, u string) (*http.Request, error)
		Final            Final
		FinalIsCall      bool
		ResultStatusCode int
	}

	testcases := map[string]testcase{
		"Request without auth token. Unauthorized status": testcase{
			ReqGen: func(method, URL string) (*http.Request, error) {
				return http.NewRequest(method, URL, nil)
			},
			Final:            Final{},
			ResultStatusCode: http.StatusUnauthorized,
		},
		"Request with auth token. OK status": testcase{
			ReqGen: func(method, URL string) (*http.Request, error) {
				r, err := http.NewRequest(method, URL, nil)
				if err != nil {
					return nil, err
				}
				r.Header.Set(string(middleware.Token), "access-token")
				return r, nil
			},
			Final: Final{
				Next: func(ctx server.Context, w http.ResponseWriter, r *http.Request) {
					v := r.Context().Value(middleware.Token)
					if v.(string) != "access-token" {
						w.WriteHeader(http.StatusUnauthorized)
						return
					}
					w.WriteHeader(http.StatusOK)
				},
			},
			FinalIsCall:      true,
			ResultStatusCode: http.StatusOK,
		},
	}

	for n, test := range testcases {
		t.Run(n, func(t *testing.T) {
			m := middleware.InitToken(test.Final.Handler)
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				m(server.Context{}, w, r)
			}))
			defer s.Close()
			client := http.Client{}
			req, err := test.ReqGen(http.MethodGet, s.URL)
			if err != nil {
				t.Fatal("cannot create request:", err)
			}
			resp, err := client.Do(req)
			if err != nil {
				t.Fatal("cannot make request:", err)
			}
			if test.ResultStatusCode != resp.StatusCode {
				t.Fatalf("invalid status code. Expect: %v. Actual: %v", test.ResultStatusCode, resp.StatusCode)
			}
			if test.FinalIsCall != test.Final.IsCall {
				t.Fatalf("error in final middleware. Expect: %v. Actual: %v", test.FinalIsCall, test.Final.IsCall)
			}
		})
	}
}
