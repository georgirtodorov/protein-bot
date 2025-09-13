package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/georgirtodorov/protein-bot/internal/handlers"
)

func TestHandleWelcome(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handlers.Welcome(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", resp.StatusCode)
	}
	expectedBody := "Welcome to Protein Bot from the Cloud!"
	if string(body) != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, string(body))
	}
}
