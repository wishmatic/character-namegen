package mcp

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.uber.org/zap"
)

func New(log *zap.Logger) (*mcp.Server, error) {
	srv := mcp.NewServer(&mcp.Implementation{
		Name:    "character-namegen",
		Version: "0.1.0",
	}, nil)

	registerTools(srv, log)

	return srv, nil
}

func registerTools(srv *mcp.Server, log *zap.Logger) {
	registerGenerateName(srv, log)
}
