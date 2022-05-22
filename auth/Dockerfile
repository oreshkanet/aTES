FROM golang:1.16-alpine

WORKDIR /auth

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /auth-app ./cmd/app/main.go

CMD [ "/auth-api" ]