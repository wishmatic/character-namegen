package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.uber.org/zap"
)

type bearerTransport struct {
	key string
}

func (b *bearerTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r = r.Clone(r.Context())
	r.Header.Set("Authorization", "Bearer "+b.key)

	return http.DefaultTransport.RoundTrip(r)
}

func connect(t *testing.T, srv *Server, transport http.RoundTripper) *mcp.ClientSession {
	t.Helper()

	httpSrv := httptest.NewServer(srv.router)
	t.Cleanup(httpSrv.Close)

	client := mcp.NewClient(&mcp.Implementation{Name: "test", Version: "0.0.0"}, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	t.Cleanup(cancel)

	cs, err := client.Connect(ctx, &mcp.StreamableClientTransport{
		Endpoint:             httpSrv.URL + "/mcp",
		HTTPClient:           &http.Client{Transport: transport},
		DisableStandaloneSSE: true,
	}, nil)
	if err != nil {
		t.Fatalf("Connect() unexpected error: %v", err)
	}

	t.Cleanup(func() { _ = cs.Close() })

	return cs
}

func TestMCPListsAndCallsGenerateName(t *testing.T) {
	srv, err := New(testConfig(), zap.NewNop())
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	cs := connect(t, srv, &bearerTransport{key: testConfig().APIKey})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tools, err := cs.ListTools(ctx, nil)
	if err != nil {
		t.Fatalf("ListTools() unexpected error: %v", err)
	}

	var found bool
	for _, tool := range tools.Tools {
		if tool.Name == "generate_name" {
			found = true
		}
	}
	if !found {
		t.Fatalf("generate_name not registered; tools = %v", tools.Tools)
	}

	res, err := cs.CallTool(ctx, &mcp.CallToolParams{
		Name:      "generate_name",
		Arguments: map[string]any{"gender": "female", "culture": "Nordic"},
	})
	if err != nil {
		t.Fatalf("CallTool() unexpected error: %v", err)
	}
	if res.IsError {
		t.Fatalf("CallTool() returned tool error: %v", res.Content)
	}

	if !strings.Contains(contentText(res.Content), "given_name") {
		t.Errorf("CallTool() result missing given_name: %v", res.Content)
	}
}

func TestMCPRejectsUnauthenticatedClient(t *testing.T) {
	srv, err := New(testConfig(), zap.NewNop())
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	httpSrv := httptest.NewServer(srv.router)
	defer httpSrv.Close()

	client := mcp.NewClient(&mcp.Implementation{Name: "test", Version: "0.0.0"}, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cs, err := client.Connect(ctx, &mcp.StreamableClientTransport{
		Endpoint:             httpSrv.URL + "/mcp",
		DisableStandaloneSSE: true,
	}, nil)
	if err == nil {
		_ = cs.Close()
		t.Fatal("Connect() expected error for unauthenticated client")
	}
}

func contentText(content []mcp.Content) string {
	var b strings.Builder
	for _, c := range content {
		if tc, ok := c.(*mcp.TextContent); ok {
			b.WriteString(tc.Text)
		}
	}

	return b.String()
}
