package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/peoxia/user-api/user"
)

type mockUserDeleter struct {
	deleteUser func(userID string) error
}

func (m mockUserDeleter) DeleteUser(userID string) error {
	return m.deleteUser(userID)
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		storage user.Deleter
		userID  string
	}
	tests := []struct {
		name           string
		args           args
		expectedStatus int
	}{
		{
			"Valid request should return 204",
			args{
				storage: mockUserDeleter{
					deleteUser: func(userID string) error {
						return nil
					},
				},
				userID: "d2a7924e-765f-4949-bc4c-219c956d0f8b",
			},
			204,
		},
		{
			"Request with malformed user ID should return 400",
			args{
				storage: mockUserDeleter{
					deleteUser: func(userID string) error {
						return nil
					},
				},
				userID: "notUUID",
			},
			400,
		},
		{
			"Database error should return 500",
			args{
				storage: mockUserDeleter{
					deleteUser: func(userID string) error {
						return fmt.Errorf("connection to database was reset")
					},
				},
				userID: "d2a7924e-765f-4949-bc4c-219c956d0f8b",
			},
			500,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/v1/users/"+tc.args.userID, nil)
			res := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/api/v1/users/{user-id}", DeleteUser(tc.args.storage))

			router.ServeHTTP(res, req)

			if got := res.Code; got != tc.expectedStatus {
				t.Errorf("Unexpected response code: got %d, exp %d", got, tc.expectedStatus)
			}
		})
	}
}
