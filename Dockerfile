FROM golang:1.27.1

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build  ./cmd/main/main.go

CMD [ "./main" ]