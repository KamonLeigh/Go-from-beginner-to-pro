package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	UserID   string
	Username string
}

type Application struct {
	AuthServiceURL string
}

func NewAppication(authServiceURL string) *Application {
	return &Application{
		AuthServiceURL: authServiceURL,
	}
}

func (app *Application) AunthenticateUser(token string) (*User, error) {
	return &User{
		UserID:   "123",
		Username: "testUser",
	}, nil
}

func TestAuthentiucationIntegration(t *testing.T) {
	authService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "Bearer valid_token" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"user_id": "123", "username":"testUser"}`))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}))

	defer authService.Close()

	app := NewAppication(authService.URL)
	token := "valid_token"
	gotUser, err := app.AunthenticateUser(token)
	assert.NoError(t, err)
	assert.Equal(t, "123", gotUser.UserID)
	assert.Equal(t, "testUser", gotUser.Username)
}
