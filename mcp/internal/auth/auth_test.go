package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"
)

const testKey = "secret-key-123"

func guarded() http.Handler {
	mw := Middleware(zap.NewNop(), testKey)

	return mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}))
}

func do(t *testing.T, h http.Handler, authHeader, sessionID string) *httptest.ResponseRecorder {
	t.Helper()

	req := httptest.NewRequest(http.MethodPost, "/mcp", nil)
	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}
	if sessionID != "" {
		req.Header.Set(sessionIDHeader, sessionID)
	}

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	return rec
}

func TestBearerAuth(t *testing.T) {
	tests := []struct {
		name       string
		authHeader string
		wantCode   int
	}{
		{name: "valid token", authHeader: "Bearer " + testKey, wantCode: http.StatusOK},
		{name: "missing header", authHeader: "", wantCode: http.StatusUnauthorized},
		{name: "wrong token", authHeader: "Bearer wrong-key", wantCode: http.StatusUnauthorized},
		{name: "empty token", authHeader: "Bearer ", wantCode: http.StatusUnauthorized},
		{name: "wrong scheme", authHeader: "Basic " + testKey, wantCode: http.StatusUnauthorized},
		{name: "no space after bearer", authHeader: "Bearer" + testKey, wantCode: http.StatusUnauthorized},
		{name: "token with trailing space", authHeader: "Bearer " + testKey + " ", wantCode: http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := do(t, guarded(), tt.authHeader, "")
			if rec.Code != tt.wantCode {
				t.Errorf("status = %d, want %d (body: %q)", rec.Code, tt.wantCode, rec.Body.String())
			}
		})
	}
}

func TestBearerAllowsSessionIDWithoutCredentials(t *testing.T) {
	rec := do(t, guarded(), "", "some-session-id")
	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d (body: %q)", rec.Code, http.StatusOK, rec.Body.String())
	}
}

func TestUnauthorizedIncludesChallenge(t *testing.T) {
	rec := do(t, guarded(), "", "")
	if rec.Header().Get("WWW-Authenticate") == "" {
		t.Error("expected WWW-Authenticate header on 401 response")
	}
}
