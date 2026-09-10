
ARG BUILDPLATFORM=linux/amd64
ARG TARGETOS=linux
ARG TARGETARCH=amd64

FROM --platform=$BUILDPLATFORM golang:1.27.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /server ./cmd/main/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /

COPY --from=builder /server /server

CMD ["/server"]