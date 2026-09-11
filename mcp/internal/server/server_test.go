package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/wishmatic/character-namegen/mcp/internal/auth"
	"github.com/wishmatic/character-namegen/mcp/internal/config"
	"go.uber.org/zap"
)

func testConfig() config.Config {
	return config.Config{
		Host:                "127.0.0.1",
		Port:                0,
		APIKey:              "secret-key-123",
		LogLevel:            "info",
		WriteTimeoutSeconds: 600,
	}
}

func TestNewRequiresAPIKey(t *testing.T) {
	cfg := testConfig()
	cfg.APIKey = ""

	if _, err := New(cfg, zap.NewNop()); !errors.Is(err, auth.ErrNoAPIKey) {
		t.Fatalf("New() error = %v, want %v", err, auth.ErrNoAPIKey)
	}
}

func TestHealthz(t *testing.T) {
	srv, err := New(testConfig(), zap.NewNop())
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	rec := httptest.NewRecorder()
	srv.router.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/healthz", nil))

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestMCPRequiresAuthorization(t *testing.T) {
	srv, err := New(testConfig(), zap.NewNop())
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	rec := httptest.NewRecorder()
	srv.router.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/mcp", strings.NewReader("{}")))

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
	}
}
