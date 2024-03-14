FROM golang:1.16-alpine

WORKDIR /auth

COPY ./auth/go.mod ./
COPY ./auth/go.sum ./
RUN go mod download

COPY ./auth ./

RUN go build -o /auth-migrator ./cmd/migrator/main.go

CMD [ "/auth-migrator" ]