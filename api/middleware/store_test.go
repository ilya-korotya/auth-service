package middleware_test

import (
	"context"
	"errors"
	"github.com/gospeak/auth-service/api/server"
	"github.com/gospeak/auth-service/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gospeak/auth-service/api/middleware"
	"github.com/gospeak/auth-service/mock"
)

func TestCacheStore(t *testing.T) {
	type testcase struct {
		Store            *mock.TokenStoreMock
		Final            Final
		FinalIsCall      bool
		ReqGen           func(u, m string) (*http.Request, error)
		ResultStatusCode int
	}

	testcases := map[string]testcase{
		"Find token in long store. Status OK": testcase{
			Store: &mock.TokenStoreMock{
				GetFunc: func(model.Filter) (*model.Token, error) {
					return &model.Token{
						ID:      "mock-id",
						Content: "user-token",
						UserID:  "user-id",
					}, nil
				},
			},
			ReqGen: func(m, u string) (*http.Request, error) {
				return http.NewRequest(m, u, nil)
			},
			ResultStatusCode: http.StatusOK,
		},
		"Don't find token in long store. Call next middleware": testcase{
			Store: &mock.TokenStoreMock{
				GetFunc: func(model.Filter) (*model.Token, error) {
					return nil, errors.New("token not find")
				},
			},

			Final: Final{
				Next: func(ctx server.Context, w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusTeapot)
					return
				},
			},
			FinalIsCall: true,
			ReqGen: func(m, u string) (*http.Request, error) {
				return http.NewRequest(m, u, nil)
			},
			ResultStatusCode: http.StatusTeapot,
		},
	}

	for n, test := range testcases {
		t.Run(n, func(t *testing.T) {
			ctx := server.Context{
				DB: &model.Store{
					Token: test.Store,
				},
			}
			m := middleware.CheckStore(test.Final.Handler)
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r = r.WithContext(context.WithValue(r.Context(), middleware.Token, "user-token"))
				m(ctx, w, r)
			}))
			defer s.Close()
			client := http.Client{}
			req, err := test.ReqGen(http.MethodGet, s.URL)
			if err != nil {
				t.Fatal("cannot create request:", err)
			}
			res, err := client.Do(req)
			if err != nil {
				t.Fatal("invalid request to server:", err)
			}
			if test.ResultStatusCode != res.StatusCode {
				t.Fatalf("invalid status code. Expect: %v. Actual: %v", test.ResultStatusCode, res.StatusCode)
			}
			if test.FinalIsCall != test.Final.IsCall {
				t.Fatalf("error in final middleware. Expect: %v. Actual: %v", test.FinalIsCall, test.Final.IsCall)
			}
		})
	}
}
