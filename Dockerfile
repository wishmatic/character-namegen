# syntax=docker/dockerfile:1

FROM golang:1.27-alpine AS build

WORKDIR /src

COPY go.mod ./
COPY mcp/go.mod mcp/go.sum ./mcp/

RUN cd mcp && go mod download

COPY . .

RUN cd mcp && CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /bin/server ./cmd/mcp

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /bin/server /server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/server"]
