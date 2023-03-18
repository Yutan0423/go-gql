FROM golang:1.20-alpine3.17 AS go-gql
RUN go install github.com/cosmtrek/air@latest
COPY ../../app /app
WORKDIR /app
CMD ["air", "-c", ".air.toml"]