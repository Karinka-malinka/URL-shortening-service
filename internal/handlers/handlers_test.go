package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/URL-shortening-service/internal/app"
)

func TestShorteningURL(t *testing.T) {
	baseURL := "http://test.com"

	handler := ShorteningURL(baseURL)

	t.Run("Should respond with status code 400 when body is empty", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/", http.NoBody)

		rr := httptest.NewRecorder()

		handler(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {

			t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, status)

		}

	})

	t.Run("Should respond with status code 400 when body has no content", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))

		rr := httptest.NewRecorder()

		handler(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {

			t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, status)

		}

	})

	t.Run("Should respond with correct short URL when sending proper URL as body", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("https://example.com"))

		rr := httptest.NewRecorder()

		handler(rr, req)

		expectedShort := baseURL + "/" + app.ShorteningURL("https://example.com")

		if status := rr.Code; status != http.StatusCreated {

			t.Errorf("Expected status code %v, got %v", http.StatusCreated, status)

		}

		if contentType := rr.Header().Get("Content-Type"); contentType != "text/plain" {

			t.Errorf("Expected Content-Type header %v, got %v", "text/plain", contentType)

		}

		expectedBody := expectedShort

		receivedBody := strings.TrimSpace(rr.Body.String())

		if expectedBody != receivedBody {

			t.Errorf("Expected body %v, got %v", expectedBody, receivedBody)

		}

	})

}
