FROM golang:1.17-alpine

ENV CGO_ENABLED=0

WORKDIR /app

COPY go.* .

RUN go mod download

COPY . .

ENTRYPOINT ["go", "test", "-v", "./...", "-coverprofile", "cover.out"]
