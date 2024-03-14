FROM golang:1.16-alpine

WORKDIR /auth

COPY ./auth/go.mod ./
COPY ./auth/go.sum ./
RUN go mod download

COPY ./auth ./
COPY ./event-registry/schemas ./schemas
ENV SCHEMA_REGISTRY_PATH="/schemas"

RUN go build -o /auth-app ./cmd/app/main.go

CMD [ "/auth-app" ]