FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./ ./

RUN go build -C ./cmd/banner-manager/ -o app

EXPOSE 8080

CMD ["./cmd/banner-manager/app"]