# syntax=docker/dockerfile:1.6

FROM golang:1.24.3 AS builder

WORKDIR /src

ENV CGO_ENABLED=0 GOOS=linux GOFLAGS=-mod=vendor

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags="-s -w" -o /out/sample_project presenter/api/main.go


FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=builder /out/sample_project /app/sample_project
COPY --from=builder /src/presenter/config /app/presenter/config
COPY --from=builder /src/languages /app/languages

EXPOSE 8000

ENTRYPOINT ["/app/sample_project"]