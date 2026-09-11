package auth

import (
	"crypto/subtle"
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

var ErrNoAPIKey = errors.New("API_KEY is required")

// sessionIDHeader is the MCP streamable-HTTP header carrying a session identifier on follow-up requests. Once a
// session is established by an authenticated request, subsequent requests address it via this header rather than
// re-sending credentials.
const sessionIDHeader = "Mcp-Session-Id"

func hasSession(r *http.Request) bool {
	return strings.TrimSpace(r.Header.Get(sessionIDHeader)) != ""
}

// Middleware enforces bearer authentication with the configured API key. Session-establishing requests must carry
// `Authorization: Bearer <API_KEY>`; follow-up requests carrying an Mcp-Session-Id are already authenticated by the
// session-establishing request.
func Middleware(log *zap.Logger, apiKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if hasSession(r) {
				next.ServeHTTP(w, r)

				return
			}

			token, reason, ok := bearerToken(r)
			if !ok || subtle.ConstantTimeCompare([]byte(token), []byte(apiKey)) != 1 {
				if ok {
					reason = "invalid token"
				}

				authFailed(log, r, reason)
				unauthorized(w)

				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func bearerToken(r *http.Request) (token, reason string, ok bool) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", "missing authorization header", false
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(header, prefix) {
		return "", "unsupported authorization scheme", false
	}

	token = strings.TrimSpace(header[len(prefix):])
	if token == "" {
		return "", "empty token", false
	}

	return token, "", true
}

func authFailed(log *zap.Logger, r *http.Request, reason string) {
	log.Warn("authentication failed",
		zap.String("reason", reason),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("request_id", middleware.GetReqID(r.Context())),
	)
}

func unauthorized(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Bearer realm="mcp"`)
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write([]byte("unauthorized"))
}
