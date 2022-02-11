package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/peoxia/user-api/user"
)

type mockUserCreatorUpdater struct {
	createOrUpdateUser func(user user.Data) (string, error)
}

func (m mockUserCreatorUpdater) CreateOrUpdateUser(user user.Data) (string, error) {
	return m.createOrUpdateUser(user)
}

func TestCreateOrUpdateUser(t *testing.T) {
	type args struct {
		storage user.CreatorUpdater
		user    user.Data
	}
	tests := []struct {
		name           string
		args           args
		expectedStatus int
	}{
		{
			"Valid data for creating a user should return 200",
			args{
				storage: mockUserCreatorUpdater{
					createOrUpdateUser: func(user user.Data) (string, error) {
						return "d2a7924e-765f-4949-bc4c-219c956d0f8b", nil
					},
				},
				user: user.Data{
					FirstName: "Alice",
					LastName:  "Bob",
					Nickname:  "AB123",
					Password:  "supersecurepassword",
					Email:     "alice@bob.com",
					Country:   "UK",
				},
			},
			200,
		},
		{
			"Database error should return 500",
			args{
				storage: mockUserCreatorUpdater{
					createOrUpdateUser: func(user user.Data) (string, error) {
						return "", fmt.Errorf("error fields cannot be empty")
					},
				},
				user: user.Data{
					Password: "supersecurepassword",
				},
			},
			500,
		},
		{
			"Valid data for updating a user should return 200",
			args{
				storage: mockUserCreatorUpdater{
					createOrUpdateUser: func(user user.Data) (string, error) {
						return "d2a7924e-765f-4949-bc4c-219c956d0f8b", nil
					},
				},
				user: user.Data{
					ID:       "d2a7924e-765f-4949-bc4c-219c956d0f8b",
					Password: "supersecurepassword",
				},
			},
			200,
		},
		{
			"Malformed data for updating a user should return 400",
			args{
				storage: mockUserCreatorUpdater{
					createOrUpdateUser: func(user user.Data) (string, error) {
						return "", nil
					},
				},
				user: user.Data{
					ID: "NOT-UUID",
				},
			},
			400,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var body bytes.Buffer
			err := json.NewEncoder(&body).Encode(tc.args.user)
			if err != nil {
				t.Errorf("Could not encode test payload: %v", err)
			}
			req := httptest.NewRequest(http.MethodPost, "/api/v1/users", &body)
			res := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/api/v1/users", CreateOrUpdateUser(tc.args.storage))

			router.ServeHTTP(res, req)

			if got := res.Code; got != tc.expectedStatus {
				t.Errorf("Unexpected response code: got %d, exp %d", got, tc.expectedStatus)
			}
		})
	}
}
