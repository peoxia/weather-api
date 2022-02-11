package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/peoxia/user-api/user"
)

type mockListGetter struct {
	getUsers func(pageID int, country string) ([]user.Data, error)
}

func (m mockListGetter) GetUsers(pageID int, country string) ([]user.Data, error) {
	return m.getUsers(pageID, country)
}

func TestGetUser(t *testing.T) {
	type args struct {
		storage user.ListGetter
		pageID  string
		country string
	}
	tests := []struct {
		name           string
		args           args
		expectedStatus int
	}{
		{
			"Valid request should return 200",
			args{
				storage: mockListGetter{
					getUsers: func(pageID int, country string) ([]user.Data, error) {
						return []user.Data{
							{
								ID:       "d2a7924e-765f-4949-bc4c-219c956d0f8b",
								Nickname: "AB123",
								Country:  "UK",
							},
						}, nil
					},
				},
				pageID:  "1",
				country: "UK",
			},
			200,
		},
		{
			"Request with missing parameters should return 400",
			args{
				storage: mockListGetter{
					getUsers: func(pageID int, country string) ([]user.Data, error) {
						return nil, nil
					},
				},
				pageID: "not-pageID",
			},
			400,
		},
		{
			"Database error should return 500",
			args{
				storage: mockListGetter{
					getUsers: func(pageID int, country string) ([]user.Data, error) {
						return nil, fmt.Errorf("error lost connection to database")
					},
				},
				pageID:  "1",
				country: "Italy",
			},
			500,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			target := fmt.Sprintf("/api/v1/users?pageId=%v&country=%s", tc.args.pageID, tc.args.country)
			req := httptest.NewRequest(http.MethodGet, target, nil)
			res := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/api/v1/users", GetUsers(tc.args.storage))

			router.ServeHTTP(res, req)

			if got := res.Code; got != tc.expectedStatus {
				t.Errorf("Unexpected response code: got %d, exp %d", got, tc.expectedStatus)
			}
		})
	}
}
